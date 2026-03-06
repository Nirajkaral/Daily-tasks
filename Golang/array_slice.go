package main

import "fmt"

func main() {

	// Array
	numbers := [4]int{10, 20, 30, 40}

	fmt.Println("Array values:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Slice
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println("Slice values:")
	for _, value := range slice {
		fmt.Println(value)
	}
}