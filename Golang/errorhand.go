package main

import (
	"fmt"
	"strconv"
)

func errorhand() {

	num, err := strconv.Atoi("abc")

	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}

	fmt.Println(num)
}