package main

import "fmt"

// Struct (like a class)
type Person struct {
	name string
	age  int
}

func strt() {

	// Object creation
	p1 := Person{
		name: "Niraj",
		age:  22,
	}

	fmt.Println("Name:", p1.name)
	fmt.Println("Age:", p1.age)
}