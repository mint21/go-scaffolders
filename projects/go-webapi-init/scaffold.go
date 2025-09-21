package main

import (
	"os"
	"path/filepath"
)

var structure = []string{
	"cmd/server",
	"internal/handler",
	"internal/router",
	"internal/service",
	"pkg/model",
}

func resolveTarget(args []string) string {
	if len(args) > 1 {
		if args[1] == "." {
			return "."
		}
		return args[1]
	}
	return "go-web-api"
}

func scaffold(root string) error {
	for _, dir := range structure {
		fullPath := filepath.Join(root, dir)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			return err
		}
		keepFile := filepath.Join(fullPath, ".keep")
		if _, err := os.Create(keepFile); err != nil {
			return err
		}
	}
	return nil
}
