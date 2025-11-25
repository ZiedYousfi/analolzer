package rofl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type RoflFile struct {
	FileBuffer     []byte
	Path           string
	MetadataOffset uint64
	Metadata       Metadata
	MetadataString string
}

func OpenRoflFile(path string) (*RoflFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// Ensure file is closed on failure
	defer file.Close()

	if _, err = file.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	needle := []byte(`{"gameLength"`)
	metadataOffset := uint64(0)

	if pos := bytes.Index(buf, needle); pos >= 0 {
		metadataOffset = uint64(pos)
		log.Printf("Metadata offset found at: %d", metadataOffset)
	} else {
		return nil, fmt.Errorf("metadata offset not found")
	}

	metadataBytes := buf[metadataOffset:]

	jsonBytes, err := extractJSON(metadataBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to locate metadata JSON: %w", err)
	}

	if metadata, err := UnmarshalMetadata(jsonBytes); err == nil {

		b, err := json.MarshalIndent(metadata, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling metadata: %v", err)
		}

		r := &RoflFile{
			FileBuffer:     buf,
			Path:           path,
			MetadataOffset: metadataOffset,
			Metadata:       metadata,
			MetadataString: string(b),
		}

		return r, nil
	}

	return nil, fmt.Errorf("error unmarshaling metadata")
}

func extractJSON(data []byte) ([]byte, error) {
	depth := 0
	inString := false
	escape := false
	start := -1

	for i, b := range data {
		if escape {
			escape = false
			continue
		}

		if b == '\\' && inString {
			escape = true
			continue
		}

		if b == '"' {
			inString = !inString
			continue
		}

		if inString {
			continue
		}

		switch b {
		case '{':
			if depth == 0 {
				start = i
			}
			depth++
		case '}':
			depth--
			if depth == 0 && start >= 0 {
				return data[start : i+1], nil
			}
		}
	}

	return nil, fmt.Errorf("metadata JSON did not close")
}
