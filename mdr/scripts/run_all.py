#!/usr/bin/env python3
"""Run all ROFL-processing scripts for a single replay."""
from __future__ import annotations

import argparse
import subprocess
import sys
from pathlib import Path

SCRIPTS_DIR = Path(__file__).resolve().parent
FIND_SCRIPT = SCRIPTS_DIR / "find_metadata_off.py"
EXTRACT_SCRIPT = SCRIPTS_DIR / "extract_metadata_schema.py"


def _run_script(script_path: Path, args: list[str]) -> None:
    command = [sys.executable, str(script_path), *args]
    try:
        subprocess.run(command, check=True)
    except subprocess.CalledProcessError as exc:
        raise SystemExit(f"{script_path.name} failed (exit {exc.returncode})") from exc


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Find metadata offset and infer metadata schema for a single ROFL file."
    )
    parser.add_argument("rofl_file", type=Path, help="Path to the .rofl replay file")
    parser.add_argument(
        "--version",
        default="15.23.726.9074",
        help="Replay version string to record alongside the metadata offset",
    )
    parser.add_argument(
        "--log-file",
        type=Path,
        help="Optional path that will be passed to the offset finder for logging",
    )
    parser.add_argument(
        "--known-offset",
        action="append",
        help="Additional metadata offsets to search for (hex or decimal).",
    )
    parser.add_argument(
        "--schema-output",
        type=Path,
        help="Optional file path where the inferred JSON schema should be written",
    )
    parser.add_argument(
        "--cast-numbers",
        action="store_true",
        help="Pass through to the schema extractor to cast numeric strings",
    )
    parser.add_argument(
        "--indent",
        type=int,
        default=2,
        help="Indentation level used when dumping the inferred schema",
    )
    parser.add_argument(
        "--skip-find",
        action="store_true",
        help="Skip the metadata offset finder and only run the extractor",
    )
    parser.add_argument(
        "--continue-on-find-failure",
        action="store_true",
        help="If the find step fails, continue with extractor instead of exiting",
    )
    return parser.parse_args()


def main() -> None:
    args = parse_args()
    rofl_path = args.rofl_file
    find_args = [str(rofl_path), "--version", args.version]
    if args.log_file:
        find_args.extend(["--log-file", str(args.log_file)])
    if args.known_offset:
        for offset in args.known_offset:
            find_args.extend(["--known-offset", offset])

    extract_args = [str(rofl_path), "--indent", str(args.indent)]
    if args.schema_output:
        extract_args.extend(["--output", str(args.schema_output)])
    if args.cast_numbers:
        extract_args.append("--cast-numbers")

    # Run the finder first (unless skipped), then run the extractor
    if not args.skip_find:
        try:
            _run_script(FIND_SCRIPT, find_args)
        except SystemExit:
            if args.continue_on_find_failure:
                print("Find step failed but continuing because --continue-on-find-failure was set")
            else:
                raise
    _run_script(EXTRACT_SCRIPT, extract_args)


if __name__ == "__main__":
    main()
