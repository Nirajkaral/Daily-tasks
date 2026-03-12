package main

import "fmt"

type Person struct {
    Name    string   
    Age     int      
    salary  float64  
    secrets string   
}

func (p Person) GetSalary() float64 {   
    return p.salary
}

func (p *Person) SetSalary(s float64) { 
    if s > 0 {
        p.salary = s
    } else {
        fmt.Println(" Salary must be positive!")
    }
}

func ULcases() {
    p := Person{Name: "Alice", Age: 25}
    p.SetSalary(50000)
    fmt.Println(p.Name)          
    fmt.Println(p.GetSalary())  
   
}