package main

import "fmt"

type Book struct {
    Title  string
    Author string
    Pages  int
}

func (b Book) DisplayInfo() {
    fmt.Printf(`
   Book Details:
   Title   : %s
   Author  : %s
   Pages   : %d
`, b.Title, b.Author, b.Pages)
}

func bookstrcut() {
    
    b1 := Book{Title: "Python Crash Course", Author: "Eric Matthes", Pages: 560}
    b2 := Book{Title: "The Go Programming Language", Author: "Alan Donovan", Pages: 380}
    b3 := Book{Title: "Clean Code", Author: "Robert Martin", Pages: 431}

 
    b1.DisplayInfo()
    b2.DisplayInfo()
    b3.DisplayInfo()
}