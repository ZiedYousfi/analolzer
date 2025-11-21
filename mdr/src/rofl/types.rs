use serde::{Deserialize, Serialize};

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
