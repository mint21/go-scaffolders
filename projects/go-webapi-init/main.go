package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go-webapi-init <target-directory>")
		fmt.Println("Pass '.' to scaffold into the current folder.")
		os.Exit(1)
	}

	root := os.Args[1]

	if err := scaffold(root); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Project structure created at:", root)
}
