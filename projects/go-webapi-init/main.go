package main

import (
	"fmt"
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

func main() {
	root := "go-web-api"
	for _, dir := range structure {
		fullPath := filepath.Join(root, dir)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			fmt.Println("Error creating directory:", err)
			continue
		}
		keepFile := filepath.Join(fullPath, ".keep")
		if _, err := os.Create(keepFile); err != nil {
			fmt.Println("Error creating .keep file:", err)
		}
	}
	fmt.Println("Project structure created at:", root)
}
