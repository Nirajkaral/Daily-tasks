package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func main() {

	jsonData := `{"Name":"Niraj","Age":22}`

	var s Student

	json.Unmarshal([]byte(jsonData), &s)

	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
}