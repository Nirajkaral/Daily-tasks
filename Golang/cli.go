package main

import (
	"fmt"
	"os"
)

func cli() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run day12_cli.go <name>")
		return
	}

	name := os.Args[1]
	fmt.Println("Hello", name)
}