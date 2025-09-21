package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitializerCreatesStructure(t *testing.T) {
	tmpDir := t.TempDir()
	root := filepath.Join(tmpDir, "go-web-api")

	for _, dir := range structure {
		fullPath := filepath.Join(root, dir)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			t.Fatalf("Failed to create directory: %v", err)
		}
		keepFile := filepath.Join(fullPath, ".keep")
		if _, err := os.Create(keepFile); err != nil {
			t.Fatalf("Failed to create .keep file: %v", err)
		}
	}

	for _, dir := range structure {
		fullPath := filepath.Join(root, dir)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			t.Errorf("Expected directory not found: %s", fullPath)
		}
		keepFile := filepath.Join(fullPath, ".keep")
		if _, err := os.Stat(keepFile); os.IsNotExist(err) {
			t.Errorf("Expected .keep file not found: %s", keepFile)
		}
	}
}
