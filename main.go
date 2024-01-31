package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

// Secret represents the structure of the secret data
type Secret struct {
	Data map[string]string `json:"data"`
}

func main() {
	err := realMain()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func realMain() error {
	overwrite := flag.Bool("overwrite", false, "Set to true to overwrite existing files")
	flag.Parse()

	// Read the input from stdin
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("io.ReadAll: %w", err)
	}

	// Unmarshal the JSON data
	var secret Secret
	err = json.Unmarshal(input, &secret)
	if err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	// Process each key in the data
	for key, encodedValue := range secret.Data {
		fileName := key
		if _, err := os.Stat(fileName); err == nil && !*overwrite {
			fmt.Printf("File %s already exists, skipping\n", fileName)
			continue
		}

		// Base64 decode the value
		decodedValue, err := base64.StdEncoding.DecodeString(encodedValue)
		if err != nil {
			return fmt.Errorf("base64.StdEncoding.DecodeString: %w", err)
		}

		// Write to a file
		err = os.WriteFile(fileName, decodedValue, 0644)
		if err != nil {
			return fmt.Errorf("os.WriteFile: %w", err)
		}
		fmt.Printf("Written to %s\n", fileName)
	}
	return nil
}
