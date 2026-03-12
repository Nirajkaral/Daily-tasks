package main

import "fmt"


type animal struct {
    Name string
}

func (a animal) Speak() string {
    return a.Name + " makes a sound"
}

func (a animal) Eat() {
    fmt.Println(a.Name, "is eating")
}

func (a animal) Sleep() {
    fmt.Println(a.Name, "is sleeping")
}


type dog struct {
    animal        
    breed  string
}


func (d dog) Speak() string {
    return d.Name + " says Woof! "
}

func (d dog) Fetch() {
    fmt.Println(d.Name, "is fetching the ball!")
}


type cat struct {
    animal
}


func (c cat) Speak() string {
    return c.Name + " says Meow! "
}

func (c cat) Purr() {
    fmt.Println(c.Name, "is purring... ")
}

func compo() {
    dog := dog{animal: animal{Name: "Buddy"}, breed: "Labrador"}
    cat := cat{animal: animal{Name: "Whiskers"}}

    fmt.Println(dog.Speak())   
    fmt.Println(cat.Speak())   
    dog.Eat()                  
    cat.Sleep()                
    dog.Fetch()                
    cat.Purr()                 
}