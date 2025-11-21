from datetime import datetime
import json
from pathlib import Path
import sys

class file:
    def __init__(self, path, known_meta_offsets, version):
        self.path = path
        self.known_meta_offsets = known_meta_offsets
        self.version = version

    def find_metadata_offset(self) -> int:
        """
        Parcourt le fichier binaire et trouve tous les endroits où un u32 little-endian
        est égal à known_value. Retourne le premier offset trouvé, ou lève une erreur.

        :param file_path: Path vers le fichier .rofl
        :param known_value: valeur entière connue (ex: offset réel du JSON)
        :return: offset (en bytes depuis le début du fichier)
        """

        with open(self.path, "rb") as f:
            data = f.read()

        for known_value in self.known_meta_offsets:
            print(f"Searching for metadata offset {known_value:#010x} in existing file {self.path}")
            cursor = 0
            data_length = len(data)

            while cursor < data_length - 4:
                value = int.from_bytes(data[cursor:cursor + 4], byteorder='little')
                if value == known_value:
                    print(f"Found metadata offset {known_value:#010x} at file offset {cursor:#06x}")
                    return cursor
                cursor += 1

        raise ValueError("Metadata offset not found")


SCRIPT_VERSION = "1.0.0"
LOG_FILE_PATH = Path(__file__).resolve().parent / "metadata_processing.log"


class ProcessingHistory:
    def __init__(self, log_path: Path):
        self.log_path = log_path
        self.entries = self._load_entries()

    def _load_entries(self) -> list[dict]:
        if not self.log_path.exists():
            return []

        entries: list[dict] = []
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

    def find_entry(self, file_path: Path, file_version: str) -> dict | None:
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


if __name__ == "__main__":
    file_instance = file(Path(
        "mdr/test/replays/EUW1-7610660427.rofl"
    ), [0x10AC67B, 0x10AC67B4], "15.23.726.9074")

    history = ProcessingHistory(LOG_FILE_PATH)
    already_processed = history.find_entry(file_instance.path, file_instance.version)

    if already_processed is not None:
        offset = already_processed["metadata_offset"]
        print(
            f"Skipping {file_instance.path} for {file_instance.version} (script v{SCRIPT_VERSION}); "
            f"metadata offset {offset:#06x} already recorded"
        )
        sys.exit(0)

    try:
        RESULT = file_instance.find_metadata_offset()
        history.record(file_instance.path, file_instance.version, RESULT)
        print(f"Metadata offset found at {RESULT:#06x}")
    except ValueError as e:
        print(e)
