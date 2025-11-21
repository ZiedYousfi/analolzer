mod rofl;

use anyhow::Result;
use std::env;
use tracing::{info, warn};

fn main() -> Result<()> {
    tracing_subscriber::fmt::init();

    let args: Vec<String> = env::args().collect();

    let input_path = if let Some(arg) = args.get(1) {
        arg
    } else {
        warn!("No input file provided");
        warn!("Using default path for testing purposes only");
        "mdr/test/replays/EUW1-7610660427.rofl"
    };

    // let output_path = if let Some(arg) = args.get(2) {
    //     arg
    // } else {
    //     warn!("No output file provided");
    //     warn!("Using default path for testing purposes only");
    //     "rofl_output.json"
    // };

    info!(path = input_path, "Reading ROFL file");
    let rofl_file = rofl::file::RoflFile::open(std::path::PathBuf::from(input_path))?;
    info!(
        metadata_offset = rofl_file.metadata_offset,
        "Found metadata offset"
    );

    info!(
        game_version = %rofl_file.metadata.game_version,
        "ROFL file game version"
    );

    Ok(())
}
