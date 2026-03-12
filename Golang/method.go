package main

import "fmt"

type Student struct {
	name string
}


func (s Student) greet() {
	fmt.Println("Hello my name is", s.name)
}

func method() {

	s1 := Student{name: "Niraj"}

	s1.greet()
}