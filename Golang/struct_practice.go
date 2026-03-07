package main

import "fmt"

type Product struct {
	name  string
	price int
}

func struct_practice() {

	p := Product{name: "Laptop", price: 50000}

	fmt.Println("Product:", p.name)
	fmt.Println("Price:", p.price)
}