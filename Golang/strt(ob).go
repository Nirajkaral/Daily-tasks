package main

import "fmt"


type person struct {
	name string
	age  int
}

func strt() {

	
	p1 := person{
		name: "Niraj",
		age:  22,
	}

	fmt.Println("Name:", p1.name)
	fmt.Println("Age:", p1.age)
}