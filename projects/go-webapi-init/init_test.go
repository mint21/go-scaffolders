package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScaffoldIntoNamedFolder(t *testing.T) {
	tmpDir := t.TempDir()
	target := filepath.Join(tmpDir, "my-api")

	if err := scaffold(target); err != nil {
		t.Fatalf("Scaffold failed: %v", err)
	}

	for _, dir := range structure {
		fullPath := filepath.Join(target, dir)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			t.Errorf("Missing directory: %s", fullPath)
		}
		if _, err := os.Stat(filepath.Join(fullPath, ".keep")); os.IsNotExist(err) {
			t.Errorf("Missing .keep file in: %s", fullPath)
		}
	}
}

func TestScaffoldIntoCurrentFolder(t *testing.T) {
	tmpDir := t.TempDir()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	if err := scaffold("."); err != nil {
		t.Fatalf("Scaffold failed: %v", err)
	}

	for _, dir := range structure {
		fullPath := filepath.Join(tmpDir, dir)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			t.Errorf("Missing directory: %s", fullPath)
		}
		if _, err := os.Stat(filepath.Join(fullPath, ".keep")); os.IsNotExist(err) {
			t.Errorf("Missing .keep file in: %s", fullPath)
		}
	}
}
