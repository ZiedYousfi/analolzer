#!/usr/bin/env python3
"""Generate a JSON Schema from the ROFL metadata stored inside a replay file.

IMPORTANT: The generated schema marks integer fields as accepting both integer
and string types because ROFL files can encode numeric values as strings.
When using quicktype.io to generate Go code from this schema, you'll need to
manually add a FlexInt64 type that can unmarshal from both integers and strings.

The statsJson field in ROFL files is a JSON-encoded string, not a direct array.
The Go code needs a custom UnmarshalJSON method to handle this two-step parsing.
"""
from __future__ import annotations

import argparse
import copy
import json
from pathlib import Path
from typing import Any

METADATA_START = b'{"gameLength"'
METADATA_END_SENTINEL = b']"}&'

# Fields that are known to be strings (not numeric)
STRING_FIELDS = {
    "WIN",
    "NAME",
    "PUUID",
    "SKIN",
    "TEAM_POSITION",
    "INDIVIDUAL_POSITION",
    "RIOT_ID_GAME_NAME",
}


def extract_metadata_json_bytes(path: Path) -> bytes:
    data: bytes = path.read_bytes()
    start = data.find(METADATA_START)
    if start < 0:
        raise ValueError(f"Metadata header not found in {path}")

    end = data.find(METADATA_END_SENTINEL, start)
    if end < 0:
        raise ValueError(f"Metadata end sentinel not found in {path}")

    end = end + len(METADATA_END_SENTINEL) - 1
    return data[start:end]


def infer_schema(value: Any, field_name: str | None = None) -> dict[str, Any]:
    if isinstance(value, dict):
        props = {name: infer_schema(val, field_name=name) for name, val in value.items()}
        return {
            "type": "object",
            "properties": props,
            "required": sorted(props.keys()),
        }

    if isinstance(value, list):
        merged_items = None
        for entry in value:
            merged_items = merge_schema(merged_items, infer_schema(entry))
        schema: dict[str, Any] = {"type": "array"}
        if merged_items is not None:
            schema["items"] = merged_items
        return schema

    if isinstance(value, bool):
        return {"type": "boolean"}
    if isinstance(value, int):
        # ROFL files can encode integers as strings, so we accept both types
        # This ensures generated Go code uses a flexible type like FlexInt64
        return {"type": ["integer", "string"]}
    if isinstance(value, float):
        return {"type": ["number", "string"]}
    if value is None:
        return {"type": "null"}

    # Check if this string field is actually a known string field
    if field_name and field_name in STRING_FIELDS:
        return {"type": "string"}

    # For other string values, check if they look like numbers
    # If so, mark them as potentially both integer and string
    if isinstance(value, str):
        # Check if it's a numeric string
        try:
            int(value)
            return {"type": ["integer", "string"]}
        except ValueError:
            pass
        try:
            float(value)
            return {"type": ["number", "string"]}
        except ValueError:
            pass

    return {"type": "string"}


def merge_schema(
    schema_a: dict[str, Any] | None, schema_b: dict[str, Any] | None
) -> dict[str, Any] | None:
    if schema_a is None:
        return copy.deepcopy(schema_b)
    if schema_b is None:
        return copy.deepcopy(schema_a)

    merged: dict[str, Any] = {}
    merged_types = _merge_types(schema_a.get("type"), schema_b.get("type"))
    merged["type"] = merged_types[0] if len(merged_types) == 1 else merged_types

    properties_a = schema_a.get("properties", {})
    properties_b = schema_b.get("properties", {})
    if properties_a or properties_b:
        keys = set(properties_a) | set(properties_b)
        props: dict[str, Any] = {}
        for key in keys:
            merged_prop = merge_schema(properties_a.get(key), properties_b.get(key))
            if merged_prop is not None:
                props[key] = merged_prop
        if props:
            merged["properties"] = props
            merged["required"] = sorted(
                set(schema_a.get("required", [])) | set(schema_b.get("required", []))
            )

    items_a = schema_a.get("items")
    items_b = schema_b.get("items")
    merged_items = merge_schema(items_a, items_b)
    if merged_items is not None:
        merged["items"] = merged_items

    return merged


def _normalize_type(value: Any) -> list[str]:
    if value is None:
        return []
    if isinstance(value, list):
        return value
    return [value]


def _merge_types(a: Any, b: Any) -> list[str]:
    normalized = _normalize_type(a) + _normalize_type(b)
    seen: list[str] = []
    for entry in normalized:
        if entry not in seen:
            seen.append(entry)
    return seen or ["object"]


def auto_cast(value: Any) -> Any:
    """Try to cast JSON string values to int/float when possible."""
    if not isinstance(value, str):
        return value

    # int ?
    if value.isdigit():
        return int(value)

    # float ?
    try:
        return float(value)
    except ValueError:
        return value


def normalize_json(value: Any) -> Any:
    """Recursively apply auto_cast to all values."""
    if isinstance(value, dict):
        return {k: normalize_json(auto_cast(v)) for k, v in value.items()}
    if isinstance(value, list):
        return [normalize_json(v) for v in value]
    return auto_cast(value)


def main() -> None:
    parser = argparse.ArgumentParser(
        description="Emit a JSON Schema inferred from the ROFL metadata block."
    )
    parser.add_argument("rofl_file", type=Path, help="Path to the .rofl replay file")
    parser.add_argument(
        "-o",
        "--output",
        type=Path,
        help="Optional path to write the generated schema",
    )
    parser.add_argument(
        "-i", "--indent", type=int, default=2, help="JSON indentation level"
    )
    parser.add_argument(
        "--cast-numbers",
        action="store_true",
        help="Cast numeric-looking strings to int/float before inferring schema",
    )

    args = parser.parse_args()
    raw_metadata = extract_metadata_json_bytes(args.rofl_file)

    try:
        metadata = json.loads(raw_metadata)
    except json.JSONDecodeError as exc:
        raise SystemExit(
            f"Failed to parse metadata JSON from {args.rofl_file}: {exc}"
        )

    # 1) Décoder statsJson si présent
    stats_raw = metadata.get("statsJson")
    if isinstance(stats_raw, str):
        try:
            stats = json.loads(stats_raw)
        except json.JSONDecodeError as exc:
            raise SystemExit(
                f"Failed to parse statsJson from {args.rofl_file}: {exc}"
            )
        # On remplace la string par la vraie structure
        metadata["statsJson"] = stats

    # 2) Optionnel : caster tous les nombres-string
    if args.cast_numbers:
        metadata = normalize_json(metadata)

    # 3) Inférer le schéma sur TOUT l'objet metadata enrichi
    schema = infer_schema(metadata)

    # Add metadata to the schema
    schema["$schema"] = "https://json-schema.org/draft/2020-12/schema"
    schema["title"] = f"{args.rofl_file.name} full metadata"
    schema["description"] = (
        "Schema for ROFL replay file metadata. IMPORTANT: The statsJson field "
        "is stored as a JSON-encoded string in the actual file and must be "
        "parsed in two steps. Additionally, numeric fields may appear as strings "
        "and should be handled with flexible parsing (e.g., FlexInt64 in Go)."
    )

    # Add description to statsJson to document it's string-encoded in the actual file
    if "properties" in schema and "statsJson" in schema["properties"]:
        schema["properties"]["statsJson"]["description"] = (
            "NOTE: In the actual ROFL file, this field is a JSON-encoded string, "
            "not a direct array. When parsing, you need to first unmarshal the "
            "outer JSON, then parse the statsJson string as JSON again. "
            "Also, numeric fields may be encoded as strings (e.g., '0' instead of 0)."
        )

    output_content = json.dumps(schema, indent=args.indent)
    if args.output:
        # Ensure the parent directory exists so writing the output file doesn't fail
        try:
            args.output.parent.mkdir(parents=True, exist_ok=True)
        except OSError:
            # If mkdir fails (e.g., permission issues), let the write raise a more informative error
            pass
        args.output.write_text(output_content, encoding="utf-8")
    else:
        print(output_content)


if __name__ == "__main__":
    main()
