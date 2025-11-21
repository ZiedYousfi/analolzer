use binrw::BinRead;
use serde::{Deserialize, Serialize};

#[derive(Debug)]
pub struct BinHeader {
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
