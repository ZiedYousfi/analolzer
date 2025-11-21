use super::error::RoflError;
use blowfish::Blowfish;
use blowfish::cipher::{BlockDecrypt, KeyInit, generic_array::GenericArray};
use byteorder::LittleEndian;

pub fn decrypt_blowfish(data: &[u8], key: &[u8]) -> Result<Vec<u8>, RoflError> {
    // Blowfish key setup
    // The key length for Blowfish can be variable (4 to 56 bytes).
    let cipher = Blowfish::<LittleEndian>::new_from_slice(key)
        .map_err(|e| RoflError::InvalidKeyLength(e.to_string()))?;

    // ECB mode decryption (implied by lack of IV and block-by-block nature in similar tools)
    // We need to process 8-byte blocks.

    let mut output = data.to_vec();
    let block_size = 8usize;

    if output.len() % block_size != 0 {
        return Err(RoflError::InvalidBlockSize);
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
            output.truncate(output.len() - padding_len);
        }
    }

    Ok(output)
}
