use anyhow::{Result, anyhow};
use base64::prelude::*;
use binrw::{BinRead, BinReaderExt};
use blowfish::Blowfish;
use blowfish::cipher::{BlockDecrypt, KeyInit, generic_array::GenericArray};
use byteorder::LittleEndian;
use flate2::read::GzDecoder;
use serde::{Deserialize, Serialize};
use std::io::{Cursor, Read, Seek, SeekFrom};

#[derive(BinRead, Debug)]
#[br(magic = b"RIOT")]
pub struct BinHeader {
    #[br(pad_before = 2)]
    pub signature: [u8; 256],
    pub header_size: u16,
    pub file_size: u32,
    pub metadata_offset: u32,
    pub metadata_size: u32,
    pub payload_header_offset: u32,
    pub payload_header_size: u32,
    pub payload_offset: u32,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Metadata {
    #[serde(rename = "gameLength")]
    pub game_length: u64,
    #[serde(rename = "gameVersion")]
    pub game_version: String,
    #[serde(rename = "lastGameChunkId")]
    pub last_game_chunk_id: u32,
    #[serde(rename = "lastKeyFrameId")]
    pub last_key_frame_id: u32,
    #[serde(rename = "statsJson")]
    pub stats_json: String,
}

#[derive(BinRead, Debug)]
pub struct PayloadHeader {
    pub match_id: u64,
    pub match_duration: u32,
    pub keyframe_count: u32,
    pub chunk_count: u32,
    pub last_chunk_id: u32,
    pub first_chunk_id: u32,
    pub keyframe_interval: u32,
    pub encryption_key_length: u16,
    #[br(count = encryption_key_length, map = |bytes: Vec<u8>| String::from_utf8_lossy(&bytes).to_string())]
    pub encryption_key: String,
}

#[derive(BinRead, Debug, Clone)]
pub struct SegmentHeader {
    pub segment_id: u32,
    pub segment_type: u8,
    pub length: u32,
    pub chunk_id: u32,
    pub offset: u32,
}

#[derive(Debug, Serialize)]
pub struct Section {
    pub time: f32,
    pub data_len: u32,
    pub type_id: u16,
    pub params: u32,
}

pub struct RoflFile {
    pub bin_header: BinHeader,
    pub metadata: Metadata,
    pub payload_header: PayloadHeader,
    pub segment_headers: Vec<SegmentHeader>,
    pub file_path: std::path::PathBuf,
}

impl RoflFile {
    pub fn open(path: impl AsRef<std::path::Path>) -> Result<Self> {
        let path = path.as_ref();
        let mut file = std::fs::File::open(path)?;

        // Read BIN Header
        let bin_header: BinHeader = file.read_le()?;

        // Read Metadata
        file.seek(SeekFrom::Start(bin_header.metadata_offset as u64))?;
        let mut metadata_bytes = vec![0u8; bin_header.metadata_size as usize];
        file.read_exact(&mut metadata_bytes)?;
        let metadata_str = String::from_utf8(metadata_bytes)?;
        let metadata: Metadata = serde_json::from_str(&metadata_str)?;

        // Read Payload Header
        file.seek(SeekFrom::Start(bin_header.payload_header_offset as u64))?;
        let payload_header: PayloadHeader = file.read_le()?;

        // Read Segment Headers
        // The payload starts with chunk count + keyframe count Segment headers
        // The payload offset in BIN header is "Payload offset from payload header start"
        // Wait, doc says: "Payload offset from payload header start" at 284.
        // And "The payload starts with chunk count + keyframe count Segment headers"

        let payload_start =
            bin_header.payload_header_offset as u64 + bin_header.payload_offset as u64;
        file.seek(SeekFrom::Start(payload_start))?;

        let total_segments = payload_header.chunk_count + payload_header.keyframe_count;
        let mut segment_headers = Vec::with_capacity(total_segments as usize);

        for _ in 0..total_segments {
            let sh: SegmentHeader = file.read_le()?;
            segment_headers.push(sh);
        }

        Ok(Self {
            bin_header,
            metadata,
            payload_header,
            segment_headers,
            file_path: path.to_path_buf(),
        })
    }

    pub fn get_segment_data(&self, index: usize) -> Result<Vec<u8>> {
        if index >= self.segment_headers.len() {
            return Err(anyhow!("Segment index out of bounds"));
        }
        let header = &self.segment_headers[index];

        // Calculate absolute offset
        // "Segment data offset (from end of segment headers)"
        // End of segment headers = Payload Start + (Total Segments * 17)
        let payload_start =
            self.bin_header.payload_header_offset as u64 + self.bin_header.payload_offset as u64;
        let segments_headers_size =
            (self.payload_header.chunk_count + self.payload_header.keyframe_count) as u64 * 17;
        let data_start = payload_start + segments_headers_size + header.offset as u64;

        let mut file = std::fs::File::open(&self.file_path)?;
        file.seek(SeekFrom::Start(data_start))?;

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

    pub fn parse_segment(&self, index: usize) -> Result<Vec<Section>> {
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
            // Update last_time for next relative calculation?
            // The doc says "relative to last section (u8, in ms)".
            // Usually this means we add to the previous timestamp.
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

fn decrypt_blowfish(data: &[u8], key: &[u8]) -> Result<Vec<u8>> {
    // Blowfish key setup
    // The key length for Blowfish can be variable (4 to 56 bytes).
    // The doc says: "Use Blowfish to decrypt the decoded encryption key with the game ID string as the key"

    let cipher = Blowfish::<LittleEndian>::new_from_slice(key)
        .map_err(|e| anyhow!("Invalid key length: {}", e))?;

    // ECB mode decryption (implied by lack of IV and block-by-block nature in similar tools)
    // We need to process 8-byte blocks.

    let mut output = data.to_vec();
    let block_size = 8usize;

    if !output.len().is_multiple_of(block_size) {
        return Err(anyhow!("Data length not multiple of block size"));
    }

    for chunk in output.chunks_mut(block_size) {
        let mut block = GenericArray::clone_from_slice(chunk);
        cipher.decrypt_block(&mut block);
        chunk.copy_from_slice(&block);
    }

    // Remove padding
    // "remove the padding whose length is provided in the last byte of the decrypted data"
    if let Some(&padding_len) = output.last() {
        let padding_len = padding_len as usize;
        if padding_len > 0 && padding_len <= output.len() {
            // Verify padding? PKCS7 usually has all bytes equal to padding_len.
            // The doc just says "length is provided in the last byte".
            output.truncate(output.len() - padding_len);
        }
    }

    Ok(output)
}
