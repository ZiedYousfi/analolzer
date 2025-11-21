mod rofl;

use anyhow::Result;
use std::env;
use std::fs::File;
use std::io::Write;

fn main() -> Result<()> {
    let args: Vec<String> = env::args().collect();
    if args.len() < 3 {
        eprintln!("Usage: {} <input_rofl> <output_json>", args[0]);
        std::process::exit(1);
    }

    let input_path = &args[1];
    let output_path = &args[2];

    println!("Reading ROFL file: {}", input_path);
    let rofl = rofl::RoflFile::open(input_path)?;

    println!("Metadata loaded.");
    println!("Match ID: {}", rofl.payload_header.match_id);
    println!("Segments: {}", rofl.segment_headers.len());

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
            Err(_e) => {
                // Don't spam stderr, just log it in the json or ignore
                // eprintln!("\nError parsing segment {}: {}", i, e);
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

    println!("Written info to {}", output_path);

    Ok(())
}
