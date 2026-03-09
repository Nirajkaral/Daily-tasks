package main

import "fmt"

// Parent struct
type Animal struct {
	name string
}

func (a Animal) Sound() {
	fmt.Println("Animal makes sound")
}

// Child struct
type Dog struct {
	Animal
}

func compo() {

	d := Dog{
		Animal{name: "Tommy"},
	}

	fmt.Println("Dog Name:", d.name)

	d.Sound()
}