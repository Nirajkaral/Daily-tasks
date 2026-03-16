package main

import "fmt"

type Student1 struct {
	name string
	age  int
}

func structs() {

	s1 := Student1{name: "Niraj", age: 22}

	fmt.Println("Name:", s1.name)
	fmt.Println("Age:", s1.age)
}
