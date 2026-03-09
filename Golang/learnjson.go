package main

import (
	"encoding/json"
	"fmt"
)

type Studentfmt struct {
	Name string
	Age  int
}

func learnjson() {

	jsonData := `{"Name":"Niraj","Age":22}`

	var s Studentfmt

	json.Unmarshal([]byte(jsonData), &s)

	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
}