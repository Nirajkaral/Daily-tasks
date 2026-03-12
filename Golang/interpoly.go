package main

import "fmt"

type Speaker interface {
    Speak() string
}

type Animal struct{ Name string }
type Dog struct{ Animal }
type Cat struct{ Animal }
type Cow struct{ Animal }


func (d Dog) Speak() string { return d.Name + " says Woof!" }
func (c Cat) Speak() string { return c.Name + " says Meow!" }
func (c Cow) Speak() string { return c.Name + " says Moo! " }

func interpoly() {
    speakers := []Speaker{
        Dog{Animal: Animal{Name: "Buddy"}},
        Cat{Animal: Animal{Name: "Whiskers"}},
        Cow{Animal: Animal{Name: "Bessie"}},
    }

    for _, s := range speakers {
        fmt.Println(s.Speak())
    }
}
