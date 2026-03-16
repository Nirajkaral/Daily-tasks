package main

import "fmt"

type Vehicle interface {
	Start()
}

type Car struct{}

type Bike struct{}

func (c Car) Start() {
	fmt.Println("Car started")
}

func (b Bike) Start() {
	fmt.Println("Bike started")
}

func abs() {

	var v Vehicle

	v = Car{}
	v.Start()

	v = Bike{}
	v.Start()
}
