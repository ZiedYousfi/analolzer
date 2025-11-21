#!/usr/bin/env python3
"""Discover the metadata offset for a ROFL replay file."""
from __future__ import annotations

import argparse
import json
import sys
from datetime import datetime
from pathlib import Path
from typing import Any, Iterable

SCRIPT_VERSION = "1.0.0"
DEFAULT_KNOWN_OFFSETS = (0x10AC67B, 0x10AC67B4)
DEFAULT_LOG_FILE = Path(__file__).resolve().parent / "metadata_processing.log"


def _parse_offset(value: str) -> int:
    return int(value, 0)


class MetadataOffsetFinder:
    def __init__(self, path: Path, known_meta_offsets: Iterable[int]) -> None:
        self.path = path
        self.known_meta_offsets = tuple(known_meta_offsets)

    def find_metadata_offset(self) -> int:
        data = self.path.read_bytes()
        data_length = len(data)

        for known_value in self.known_meta_offsets:
            print(
                f"Searching for metadata offset {known_value:#010x} in {self.path}"
            )
            cursor = 0
            while cursor < data_length - 3:
                value = int.from_bytes(data[cursor : cursor + 4], byteorder="little")
                if value == known_value:
                    print(
                        f"Found metadata offset {known_value:#010x} at file offset {cursor:#06x}"
                    )
                    return cursor
                cursor += 1

        raise ValueError("Metadata offset not found")


class ProcessingHistory:
    def __init__(self, log_path: Path) -> None:
        self.log_path = log_path
        self.entries = self._load_entries()

    def _load_entries(self) -> list[dict[str, Any]]:
        if not self.log_path.exists():
            return []

        entries: list[dict[str, Any]] = []
        with self.log_path.open("r", encoding="utf-8") as log_file:
            for raw_line in log_file:
                line = raw_line.strip()
                if not line:
                    continue
                try:
                    entries.append(json.loads(line))
                except json.JSONDecodeError:
                    continue
        return entries

    def find_entry(self, file_path: Path, file_version: str) -> dict[str, Any] | None:
        normalized_path = str(file_path)
        for entry in self.entries:
            if (
                entry.get("file_path") == normalized_path
                and entry.get("file_version") == file_version
                and entry.get("script_version") == SCRIPT_VERSION
            ):
                return entry
        return None

    def record(self, file_path: Path, file_version: str, metadata_offset: int) -> None:
        self.log_path.parent.mkdir(parents=True, exist_ok=True)
        entry = {
            "processed_at": datetime.utcnow().isoformat() + "Z",
            "script_version": SCRIPT_VERSION,
            "file_path": str(file_path),
            "file_version": file_version,
            "metadata_offset": metadata_offset,
        }

        with self.log_path.open("a", encoding="utf-8") as log_file:
            json.dump(entry, log_file, separators=(",", ":"))
            log_file.write("\n")

        self.entries.append(entry)


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Find the metadata offset for a ROFL replay file."
    )
    parser.add_argument("rofl_file", type=Path, help="Path to the .rofl replay file")
    parser.add_argument(
        "--version",
        default="15.23.726.9074",
        help="Replay version to associate with the discovery",
    )
    parser.add_argument(
        "--log-file",
        type=Path,
        default=DEFAULT_LOG_FILE,
        help="Path to the processing history log",
    )
    parser.add_argument(
        "--known-offset",
        type=_parse_offset,
        action="append",
        help="Known metadata offsets encoded as integers (hex or decimal); defaults are used when omitted",
    )
    return parser.parse_args()


def main() -> None:
    args = parse_args()
    known_offsets = args.known_offset or list(DEFAULT_KNOWN_OFFSETS)

    finder = MetadataOffsetFinder(args.rofl_file, known_offsets)
    history = ProcessingHistory(args.log_file)

    processed_entry = history.find_entry(args.rofl_file, args.version)
    if processed_entry is not None:
        offset = processed_entry["metadata_offset"]
        print(
            f"Skipping {args.rofl_file} for {args.version} (script v{SCRIPT_VERSION}); "
            f"metadata offset {offset:#06x} already recorded"
        )
        sys.exit(0)

    try:
        result = finder.find_metadata_offset()
    except ValueError as exc:
        print(exc)
        sys.exit(1)

    history.record(args.rofl_file, args.version, result)
    print(f"Metadata offset found at {result:#06x}")


if __name__ == "__main__":
    main()
