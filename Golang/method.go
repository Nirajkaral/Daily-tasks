package main

import "fmt"

type Student struct {
	name string
}

// Method
func (s Student) greet() {
	fmt.Println("Hello my name is", s.name)
}

func main() {

	s1 := Student{name: "Neeraj"}

	s1.greet()
}