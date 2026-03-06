package main

import "fmt"

func fibonacci() {

	a := 0
	b := 1

	fmt.Println("Fibonacci Series:")

	for i := 0; i < 10; i++ {
		fmt.Println(a)
		next := a + b
		a = b
		b = next
	}
}