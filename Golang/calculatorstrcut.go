package main

import "fmt"


type Calculator struct {
    History []string
}


func (c Calculator) Add(a, b float64) float64 {
    return a + b
}

func (c Calculator) Subtract(a, b float64) float64 {
    return a - b
}

func (c Calculator) Multiply(a, b float64) float64 {
    return a * b
}


func (c *Calculator) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf(" cannot divide by zero")
    }
    result := a / b
    c.History = append(c.History, fmt.Sprintf("%.2f / %.2f = %.2f", a, b, result))
    return result, nil
}

func calculatorstrcut() {
    calc := Calculator{}

    fmt.Println(calc.Add(10, 5))        
    fmt.Println(calc.Subtract(10, 3))   
    fmt.Println(calc.Multiply(4, 6))    

    result, err := calc.Divide(20, 4)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)             
    }

    _, err = calc.Divide(10, 0)
    if err != nil {
        fmt.Println(err)                
    }
}