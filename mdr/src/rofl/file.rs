use super::error::RoflError;
use super::types::Metadata;

use std::fs::File;
use std::io::{Read, Seek, SeekFrom};
use std::path::PathBuf;

pub struct RoflFile {
    pub metadata_offset: u64,
    pub metadata: Metadata,
    pub file: File,
}

impl RoflFile {
    pub fn open(path: PathBuf) -> Result<Self, RoflError> {
        let file = File::open(&path)?;

        let mut rofl_file = RoflFile {
            metadata_offset: 0,
            metadata: Metadata {
                game_length: 0,
                last_game_chunk_id: 0,
                last_key_frame_id: 0,
                stats_json: Vec::new(),
            },
            file,
        };

        rofl_file.find_metadata_offset_raw()?;
        rofl_file.read_metadata()?;

        Ok(rofl_file)
    }

    pub fn find_metadata_offset_raw(&mut self) -> Result<(), RoflError> {
        let mut buf = Vec::new();
        let file = &mut self.file;
        file.seek(SeekFrom::Start(0))?;
        file.read_to_end(&mut buf)?;

        let needle = br#"{"gameLength"#;

        if let Some(pos) = buf.windows(needle.len()).position(|w| w == needle) {
            self.metadata_offset = pos as u64;
            tracing::info!(
                metadata_offset = self.metadata_offset,
                "Found metadata offset"
            );
            Ok(())
        } else {
            Err(RoflError::InvalidHeader(
                "Metadata offset not found".to_string(),
            ))
        }
    }

    pub fn read_metadata(&mut self) -> Result<(), RoflError> {
        // Relis tout le fichier en mémoire
        let mut buf = Vec::new();
        self.file.seek(SeekFrom::Start(0))?;
        self.file.read_to_end(&mut buf)?;

        let start = self.metadata_offset as usize;
        let slice = &buf[start..];

        let needle = br#"]"}&"#;
        let end = if let Some(pos) = slice.windows(needle.len()).position(|w| w == needle) {
            start + pos + needle.len() - 1
        } else {
            buf.len()
        };

        let json_slice = &buf[start..end];

        let metadata: Metadata = serde_json::from_slice(json_slice)
            .map_err(|e| RoflError::InvalidHeader(format!("Failed to parse metadata JSON: {e}")))?;

        tracing::debug!(?metadata, "Parsed metadata");

        self.metadata = metadata;
        Ok(())
    }
}
