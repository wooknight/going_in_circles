package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//go:embed replace1.txt
var replace1 string

//go:embed replace2.txt
var replace2 string

//go:embed replace3.txt
var replace3 string

//go:embed replace4.txt
var replace4 string

//go:embed replace5.txt
var replace5 string

// processFile reads a file, replaces the dataview block if found, and writes the file back.
func processFile(path string) error {
	// Read the file contents.
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	// Convert file content to a string.
	content := string(data)

	// Replace all occurrences of the dataview block.
	replacedContent := strings.ReplaceAll(content, replace1, "")
	replacedContent = strings.ReplaceAll(replacedContent, replace2, "")
	replacedContent = strings.ReplaceAll(replacedContent, replace3, "")
	replacedContent = strings.ReplaceAll(replacedContent, replace4, "")
	replacedContent = strings.ReplaceAll(replacedContent, replace5, "")
	// If no changes were made, no need to rewrite the file.
	if replacedContent == content {
		return nil
	}

	// Write the updated content back to the file.
	// Using 0644 for file permissions; adjust as needed.
	err = os.WriteFile(path, []byte(replacedContent), 0644)
	if err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	fmt.Printf("Updated file: %s\n", path)
	return nil
}

func readRecursively(root string, call func(string) error) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %q: %v\n", path, err)

			return err
		}
		if !info.IsDir() {
			call(path)
		}
		return nil
	})
	return err
}

func main() {
	root := flag.String("root", "/Users/ramesh/Documents/obsidian/obsidian", "The directory to start recursing")
	flag.Parse()

	readRecursively(*root, processFile)
}
