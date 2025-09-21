package main

import (
	"fmt"
	"os"
)

func main() {
	root := resolveTarget(os.Args)
	if err := scaffold(root); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Project structure created at:", root)
}
