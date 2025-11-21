use thiserror::Error;

#[derive(Error, Debug)]
pub enum RoflError {
    #[error("IO error: {0}")]
    Io(#[from] std::io::Error),

    #[error("JSON error: {0}")]
    Json(#[from] serde_json::Error),

    #[error("Base64 decode error: {0}")]
    Base64(#[from] base64::DecodeError),

    #[error("BinRw error: {0}")]
    BinRw(#[from] binrw::Error),

    #[error("UTF-8 error: {0}")]
    Utf8(#[from] std::string::FromUtf8Error),

    #[error("Invalid key length: {0}")]
    InvalidKeyLength(String),

    #[error("Data length not multiple of block size")]
    InvalidBlockSize,

    #[error("Segment index out of bounds")]
    SegmentIndexOutOfBounds,

    #[error("Decryption error: {0}")]
    DecryptionError(String),
    #[error("Invalid file header or offsets: {0}")]
    InvalidHeader(String),
}
