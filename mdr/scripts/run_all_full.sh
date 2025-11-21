#!/usr/bin/env bash
# run_all_full.sh — Convenience wrapper that demonstrates running run_all.py with all flags
# Usage: ./run_all_full.sh [path-to-replay]

set -euo pipefail

SCRIPTS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PYTHON=${PYTHON:-python3}

ROFL_PATH=${1:-$SCRIPTS_DIR/../test/replays/EUW1-7610660427.rofl}
VERSION=${VERSION:-15.23.726.9074}
LOG_FILE=${LOG_FILE:-$SCRIPTS_DIR/metadata_processing.log}
SCHEMA_OUTPUT=${SCHEMA_OUTPUT:-$SCRIPTS_DIR/../analysis_tmp/replay-schema.json}
INDENT=${INDENT:-2}
CAST_NUMBERS=${CAST_NUMBERS:-1}
SKIP_FIND=${SKIP_FIND:-0}
CONTINUE_ON_FIND_FAILURE=${CONTINUE_ON_FIND_FAILURE:-1}

CMD=("$PYTHON" "$SCRIPTS_DIR/run_all.py" "$ROFL_PATH" --version "$VERSION" --log-file "$LOG_FILE")

if [ -n "${KNOWN_OFFSETS:-}" ]; then
  # If the environment provides a KNOWN_OFFSETS space-separated list, use it.
  for off in ${KNOWN_OFFSETS}; do
    CMD+=(--known-offset "$off")
  done
else
  # defaults
  for off in 0x10AC67B 0x10AC67B4; do
    CMD+=(--known-offset "$off")
  done
fi

CMD+=(--schema-output "$SCHEMA_OUTPUT" --indent "$INDENT")
if [ "$CAST_NUMBERS" -eq 1 ]; then
  CMD+=(--cast-numbers)
fi

if [ "$SKIP_FIND" -eq 1 ]; then
  CMD+=(--skip-find)
fi
if [ "$CONTINUE_ON_FIND_FAILURE" -eq 1 ]; then
  CMD+=(--continue-on-find-failure)
fi

echo "Executing: ${CMD[*]}"

# Ensure the parent directory for schema output exists so Python can write the file
SCHEMA_OUTPUT_DIR="$(dirname "$SCHEMA_OUTPUT")"
mkdir -p "$SCHEMA_OUTPUT_DIR"

# Run the command and check exit status
if "${CMD[@]}"; then
  echo "Metadata extraction and schema generation completed successfully."
  echo "Schema output written to: $SCHEMA_OUTPUT"
else
  echo "An error occurred during metadata extraction or schema generation."
  exit 1
fi
