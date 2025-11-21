import json
import re
from pathlib import Path

schema_path = Path('mdr/analysis_tmp/replay-schema.json')
schema = json.loads(schema_path.read_text())
props = schema['properties']['statsJson']['items']['properties']
used = {}
keywords = {
    'as','break','const','continue','crate','else','enum','extern','false','fn','for','if','impl','in','let','loop','match','mod','move','mut','pub','ref','return','self','Self','static','struct','super','trait','true','type','unsafe','use','where','while','async','await','dyn','abstract','become','box','do','final','macro','override','priv','try','typeof','unsized','virtual','yield'
}

def snakecase(name: str) -> str:
    s = name.replace('-', '_')
    s = re.sub('(.)([A-Z][a-z]+)', r'\1_\2', s)
    s = re.sub('([a-z0-9])([A-Z])', r'\1_\2', s)
    s = s.replace('__', '_')
    return s


def sanitize(name):
    s = snakecase(name)
    s = re.sub('[^A-Za-z0-9]+', '_', s)
    s = s.lower().strip('_')
    if not s:
        s = 'field'
    if s[0].isdigit():
        s = 'field_' + s
    if s in keywords:
        s = s + '_field'
    base = s
    counter = 1
    while s in used:
        s = f"{base}_{counter}"
        counter += 1
    used[s] = name
    return s

def field_type(info):
    t = info.get('type')
    if isinstance(t, list):
        types = set(t)
        if types <= {'string', 'integer'}:
            return 'Option<serde_json::Value>'
        if types <= {'string'}:
            return 'Option<String>'
        if types <= {'integer'}:
            return 'Option<i64>'
        return 'Option<serde_json::Value>'
    if t == 'integer':
        return 'Option<i64>'
    if t == 'string':
        return 'Option<String>'
    return 'Option<serde_json::Value>'

lines = [
    'use serde::{Deserialize, Serialize};',
    '',
    '#[derive(Debug, Serialize, Deserialize)]',
    'pub struct StatsJsonEntry {',
]
for name, info in props.items():
    field = sanitize(name)
    ty = field_type(info)
    lines.append(f'    #[serde(rename = "{name}")]')
    lines.append(f'    pub {field}: {ty},')
lines.append('}')
lines.append('')
lines.extend([
    '#[derive(Debug, Serialize, Deserialize)]',
    'pub struct Metadata {',
    '    #[serde(rename = "gameLength")]',
    '    pub game_length: u64,',
    '    #[serde(rename = "lastGameChunkId")]',
    '    pub last_game_chunk_id: u32,',
    '    #[serde(rename = "lastKeyFrameId")]',
    '    pub last_key_frame_id: u32,',
    '    #[serde(rename = "statsJson")]',
    '    pub stats_json: Vec<StatsJsonEntry>,',
    '}',
])
print('\n'.join(lines))
