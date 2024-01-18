// main.go
package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	// WEPKeySize256Bits is the size of the 256-bit WEP key in bytes
	WEPKeySize256Bits = 16
)

func generateWEPKey() (string, error) {
	keyBytes := make([]byte, WEPKeySize256Bits)
	_, err := rand.Read(keyBytes)
	if err != nil {
		return "", err
	}
	upperHexKey := strings.ToUpper(hex.EncodeToString(keyBytes))

	return upperHexKey, nil
}

func main() {
	// Define flags
	outputFile := flag.String("output", "", "Output file for the WEP key")

	// Parse command line arguments
	flag.Parse()

	// Generate the WEP key
	wepKey, err := generateWEPKey()
	if err != nil {
		fmt.Println("Error generating WEP key:", err)
		os.Exit(1)
	}

	// Print or save the WEP key based on the provided output file
	if *outputFile == "" {
		fmt.Println("Generated WEP Key:", wepKey)
	} else {
		err := saveKeyToFile(*outputFile, wepKey)
		if err != nil {
			fmt.Println("Error saving WEP key to file:", err)
			os.Exit(1)
		}
		fmt.Println("WEP Key saved to", *outputFile)
	}
}

func saveKeyToFile(filename, key string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(key)
	return err
}
