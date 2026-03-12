package main

import "fmt"

type Counter struct {
    Count int
}


func (c Counter) ShowCount() {
    fmt.Println("Count:", c.Count)
}


func (c *Counter) Increment() {
    c.Count++             
}

func (c *Counter) Reset() {
    c.Count = 0
}

func reciver() {
    c := Counter{Count: 0}
    c.Increment()
    c.Increment()
    c.Increment()
    c.ShowCount()   
    c.Reset()
    c.ShowCount()   
}