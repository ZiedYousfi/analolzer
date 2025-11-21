#!/usr/bin/env python3
"""Generate a JSON Schema from the ROFL metadata stored inside a replay file."""
from __future__ import annotations

import argparse
import copy
import json
from pathlib import Path
from typing import Any

METADATA_START = b'{"gameLength"'
METADATA_END_SENTINEL = b']"}&'


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


def infer_schema(value: Any) -> dict[str, Any]:
    if isinstance(value, dict):
        props = {name: infer_schema(val) for name, val in value.items()}
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
        return {"type": "integer"}
    if isinstance(value, float):
        return {"type": "number"}
    if value is None:
        return {"type": "null"}

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
    schema = {
        "$schema": "https://json-schema.org/draft/2020-12/schema",
        "title": f"{args.rofl_file.name} full metadata",
        **schema,
    }

    output_content = json.dumps(schema, indent=args.indent)
    if args.output:
        args.output.write_text(output_content, encoding="utf-8")
    else:
        print(output_content)


if __name__ == "__main__":
    main()
