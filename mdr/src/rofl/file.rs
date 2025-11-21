use base64::prelude::*;
use binrw::BinReaderExt;
use flate2::read::GzDecoder;
use std::io::{Cursor, Read, Seek, SeekFrom};
use tracing::{debug, info, instrument};

use super::crypto::decrypt_blowfish;
use super::error::RoflError;
use super::types::{BinHeader, Metadata, PayloadHeader, Section, SegmentHeader};

pub struct RoflFile {
    pub bin_header: BinHeader,
    pub metadata: Metadata,
    pub payload_header: PayloadHeader,
    pub segment_headers: Vec<SegmentHeader>,
    pub file_path: std::path::PathBuf,
}

impl RoflFile {
    #[instrument(skip(path), fields(path = %path.as_ref().display()))]
    pub fn open(path: impl AsRef<std::path::Path>) -> Result<Self, RoflError> {
        let path = path.as_ref();
        debug!("Opening ROFL file");
        let mut file = std::fs::File::open(path)?;
        let file_len = file.metadata()?.len();

        // Read BIN Header
        let bin_header: BinHeader = file.read_le()?;
        debug!(?bin_header, "Read BIN Header");

        // Read Metadata
        file.seek(SeekFrom::Start(bin_header.metadata_offset as u64))?;
        if bin_header.metadata_offset as u64 + bin_header.metadata_size as u64 > file_len {
            return Err(RoflError::InvalidHeader(format!(
                "metadata range exceeds file size: offset {} size {} file_len {}",
                bin_header.metadata_offset, bin_header.metadata_size, file_len
            )));
        }
        let mut metadata_bytes = vec![0u8; bin_header.metadata_size as usize];
        file.read_exact(&mut metadata_bytes)?;
        let metadata_str = String::from_utf8(metadata_bytes)?;
        let metadata: Metadata = serde_json::from_str(&metadata_str)?;
        debug!(?metadata, "Read Metadata");

        // Read Payload Header
        file.seek(SeekFrom::Start(bin_header.payload_header_offset as u64))?;
        if bin_header.payload_header_offset as u64 + bin_header.payload_header_size as u64
            > file_len
        {
            return Err(RoflError::InvalidHeader(format!(
                "payload header range exceeds file size: offset {} size {} file_len {}",
                bin_header.payload_header_offset, bin_header.payload_header_size, file_len
            )));
        }
        let payload_header: PayloadHeader = file.read_le()?;
        debug!(?payload_header, "Read Payload Header");

        // Read Segment Headers
        let payload_start =
            bin_header.payload_header_offset as u64 + bin_header.payload_offset as u64;
        file.seek(SeekFrom::Start(payload_start))?;

        let total_segments = payload_header.chunk_count + payload_header.keyframe_count;
        let mut segment_headers = Vec::with_capacity(total_segments as usize);

        // Ensure the segment headers block is within file bounds
        let segments_headers_size = (total_segments as u64)
            .checked_mul(17)
            .ok_or_else(|| RoflError::InvalidHeader("segment headers size overflow".into()))?;
        if payload_start + segments_headers_size > file_len {
            return Err(RoflError::InvalidHeader(format!(
                "segment headers exceed file size: payload_start {} headers_size {} file_len {}",
                payload_start, segments_headers_size, file_len
            )));
        }

        for _ in 0..total_segments {
            let sh: SegmentHeader = file.read_le()?;
            segment_headers.push(sh);
        }
        info!(total_segments, "Loaded segment headers");

        Ok(Self {
            bin_header,
            metadata,
            payload_header,
            segment_headers,
            file_path: path.to_path_buf(),
        })
    }

    #[instrument(skip(self))]
    pub fn get_segment_data(&self, index: usize) -> Result<Vec<u8>, RoflError> {
        if index >= self.segment_headers.len() {
            return Err(RoflError::SegmentIndexOutOfBounds);
        }
        let header = &self.segment_headers[index];

        // Calculate absolute offset
        let payload_start =
            self.bin_header.payload_header_offset as u64 + self.bin_header.payload_offset as u64;
        let segments_headers_size =
            (self.payload_header.chunk_count + self.payload_header.keyframe_count) as u64 * 17;
        let data_start = payload_start + segments_headers_size + header.offset as u64;

        let mut file = std::fs::File::open(&self.file_path)?;
        let file_len = file.metadata()?.len();
        file.seek(SeekFrom::Start(data_start))?;

        // Validate the read won't go beyond the file
        if data_start + header.length as u64 > file_len {
            return Err(RoflError::InvalidHeader(format!(
                "segment {} data exceeds file size: data_start {} length {} file_len {}",
                header.segment_id, data_start, header.length, file_len
            )));
        }

        let mut encrypted_data = vec![0u8; header.length as usize];
        file.read_exact(&mut encrypted_data)?;

        // Decryption
        let game_id_str = self.payload_header.match_id.to_string();
        let key_base64 = &self.payload_header.encryption_key;
        let raw_key = BASE64_STANDARD.decode(key_base64)?;

        let chunk_key = decrypt_blowfish(&raw_key, game_id_str.as_bytes())?;
        let decrypted_data = decrypt_blowfish(&encrypted_data, &chunk_key)?;

        // Decompression
        let mut decoder = GzDecoder::new(&decrypted_data[..]);
        let mut decompressed_data = Vec::new();
        decoder.read_to_end(&mut decompressed_data)?;

        Ok(decompressed_data)
    }

    #[instrument(skip(self))]
    pub fn parse_segment(&self, index: usize) -> Result<Vec<Section>, RoflError> {
        let data = self.get_segment_data(index)?;
        let mut cursor = Cursor::new(data);
        let mut sections = Vec::new();
        let mut last_type_id = 0u16;
        let mut last_time = 0.0f32;

        while cursor.position() < cursor.get_ref().len() as u64 {
            // Read configuration byte
            let mut buf = [0u8; 1];
            if cursor.read(&mut buf)? == 0 {
                break;
            }
            let h = buf[0];

            // Game time
            let time = if h & 0x80 == 0 {
                let val: f32 = cursor.read_le()?;
                val
            } else {
                let val: u8 = cursor.read_le()?;
                last_time + (val as f32 / 1000.0)
            };
            last_time = time;

            // Data Length
            let data_len = if h & 0x10 == 0 {
                cursor.read_le::<u32>()?
            } else {
                cursor.read_le::<u8>()? as u32
            };

            // Type
            let type_id = if h & 0x40 == 0 {
                let val: u16 = cursor.read_le()?;
                last_type_id = val;
                val
            } else {
                last_type_id
            };

            // Parameters
            let params = if h & 0x20 == 0 {
                cursor.read_le::<u32>()?
            } else {
                cursor.read_le::<u8>()? as u32
            };

            // Skip data
            if data_len > 0 {
                cursor.seek(SeekFrom::Current(data_len as i64))?;
            }

            sections.push(Section {
                time,
                data_len,
                type_id,
                params,
            });
        }
        Ok(sections)
    }
}
