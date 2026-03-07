package main

import "fmt"

type Student struct {
	name string
	age  int
}

func structs() {

	s1 := Student{name: "Neeraj", age: 22}

	fmt.Println("Name:", s1.name)
	fmt.Println("Age:", s1.age)
}