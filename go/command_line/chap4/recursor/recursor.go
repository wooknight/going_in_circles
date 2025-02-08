package main

import (
	"embed"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//go:embed rules/*.txt
var rulesFS embed.FS

func getRules() ([]string, error) {
	var rules []string
	entries, err := rulesFS.ReadDir("rules")
	if err != nil {
		return nil, fmt.Errorf("reading rules: %w", err)
	}
	for _, entry := range entries {
		rule, err := rulesFS.ReadFile("rules/" + entry.Name())
		if err != nil {
			return nil, fmt.Errorf("reading rule %s: %w", entry.Name(), err)
		}
		rules = append(rules, string(rule))
	}
	return rules, nil
}

// processFile reads a file, replaces the dataview block if found, and writes the file back.
func processFile(path string) error {
	// Read the file contents.
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}
	// Convert file content to a string.
	content := string(data)
	rules, err := getRules()
	if err != nil {
		return fmt.Errorf("getting rules: %w", err)
	}
	replacedContent := content
	// Replace all occurrences of the dataview block.
	for _, rule := range rules {
		replacedContent = strings.ReplaceAll(replacedContent, rule, "")
	}
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
	root := flag.String("root", "/Users/ramesh/Documents/obsidian/obsidian/2-areas/journal/daily", "The directory to start recursing")
	flag.Parse()

	readRecursively(*root, processFile)
}
