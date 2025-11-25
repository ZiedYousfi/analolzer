package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ZiedYousfi/analolzer/mdr/rofl"
)

/// Example usage of the rofl package to open a ROFL file
/// Make sure to replace the file path with an actual ROFL file path on your system

func main() {
	roflFile, err := rofl.OpenRoflFile("/Users/ziedyousfi/code/analolzer/mdr/test/replays/EUW1-7610660427.rofl")
	if err != nil {
		log.Fatalf("Error opening ROFL file: %v", err)
	}

	fmt.Printf("ROFL file opened successfully: %s\n", roflFile.Path)
	fmt.Printf("Metadata offset: %d\n", roflFile.MetadataOffset)

	f, err := os.Create("metadata-extracted.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	l, err := f.WriteString(roflFile.MetadataString)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
}
