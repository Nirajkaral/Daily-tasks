package main

import "fmt"

func main() {

	var num int
	fmt.Print("Enter number: ")
	fmt.Scan(&num)

	fact := 1

	for i := 1; i <= num; i++ {
		fact = fact * i
	}

	fmt.Println("Factorial =", fact)
}