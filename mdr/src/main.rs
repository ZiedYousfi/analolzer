mod rofl;

use anyhow::Result;
use std::env;
use std::fs::File;
use std::io::Write;
use tracing::{info, error, warn};

fn main() -> Result<()> {
    tracing_subscriber::fmt::init();

    let args: Vec<String> = env::args().collect();
    if args.len() < 3 {
        error!("Usage: {} <input_rofl> <output_json>", args[0]);
        std::process::exit(1);
    }

    let input_path = &args[1];
    let output_path = &args[2];

    info!(path = input_path, "Reading ROFL file");
    let rofl = rofl::RoflFile::open(input_path)?;

    info!("Metadata loaded");
    info!(match_id = rofl.payload_header.match_id, "Match ID");
    info!(segments = rofl.segment_headers.len(), "Segments count");

    let mut segments_info = Vec::new();

    for (i, header) in rofl.segment_headers.iter().enumerate() {
        print!(
            "\rProcessing segment {}/{}",
            i + 1,
            rofl.segment_headers.len()
        );
        std::io::stdout().flush()?;

        // Testing only but remove this musle on clippy when done
        #[allow(clippy::manual_ok_err)]
        let sections = match rofl.parse_segment(i) {
            Ok(s) => Some(s),
            Err(e) => {
                warn!(segment_index = i, error = ?e, "Error parsing segment");
                None
            }
        };

        segments_info.push(serde_json::json!({
            "id": header.segment_id,
            "type": if header.segment_type == 1 { "Chunk" } else { "Keyframe" },
            "length": header.length,
            "sections": sections
        }));
    }
    println!();

    let output_data = serde_json::json!({
        "metadata": rofl.metadata,
        "payload_header": {
            "match_id": rofl.payload_header.match_id,
            "match_duration": rofl.payload_header.match_duration,
            "keyframe_count": rofl.payload_header.keyframe_count,
            "chunk_count": rofl.payload_header.chunk_count,
            "encryption_key": rofl.payload_header.encryption_key,
        },
        "segments": segments_info
    });

    let mut output_file = File::create(output_path)?;
    serde_json::to_writer_pretty(&mut output_file, &output_data)?;

    info!(path = output_path, "Written info to file");

    Ok(())
}
