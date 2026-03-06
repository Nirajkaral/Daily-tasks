package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func function() {

	result := add(5, 7)

	fmt.Println("Sum =", result)
}