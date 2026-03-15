package main

import (
    "fmt"
    "strings"
)

type LibraryItem interface {
    ItemType() string
    DisplayInfo()
    Borrow() bool
    ReturnItem()
    IsAvailable() bool
}

type BaseItem struct {
    Title     string
    Author    string
    ItemID    string
    available bool    
}

func (b *BaseItem) IsAvailable() bool { return b.available }

func (b *BaseItem) Borrow() bool {
    if !b.available {
        fmt.Printf(" '%s' is not available!\n", b.Title)
        return false
    }
    b.available = false
    fmt.Printf(" '%s' borrowed successfully!\n", b.Title)
    return true
}

func (b *BaseItem) ReturnItem() {
    b.available = true
    fmt.Printf(" '%s' returned successfully!\n", b.Title)
}


type BOOK struct {
    BaseItem
    Pages int
}

type Magazine struct {
    BaseItem
    Issue string
}

type DVD struct {
    BaseItem
    Duration int
}


func (b BOOK) ItemType() string { return " Book" }
func (b BOOK) DisplayInfo() {
    status := " Available"
    if !b.IsAvailable() {
        status = " Borrowed"
    }
    fmt.Printf("\n%s | %s | by %s | %d pages | %s\n",
        b.ItemType(), b.Title, b.Author, b.Pages, status)
}


func (m Magazine) ItemType() string { return " Magazine" }
func (m Magazine) DisplayInfo() {
    status := " Available"
    if !m.IsAvailable() {
        status = " Borrowed"
    }
    fmt.Printf("\n%s | %s | Issue: %s | %s\n",
        m.ItemType(), m.Title, m.Issue, status)
}


func (d DVD) ItemType() string { return " DVD" }
func (d DVD) DisplayInfo() {
    status := " Available"
    if !d.IsAvailable() {
        status = " Borrowed"
    }
    fmt.Printf("\n%s | %s | %d mins | %s\n",
        d.ItemType(), d.Title, d.Duration, status)
}


type Library struct {
    Name  string
    items []LibraryItem  
}

func (l *Library) AddItem(item LibraryItem) {
    l.items = append(l.items, item)
}

func (l *Library) ShowAll() {
    fmt.Printf("\n %s — All Items:\n", l.Name)
    fmt.Println("=" * 40)
    for _, item := range l.items {
        item.DisplayInfo()              
    }
}

func (l *Library) Search(title string) LibraryItem {
    for _, item := range l.items {
        if strings.Contains(
            strings.ToLower(fmt.Sprintf("%v", item)),
            strings.ToLower(title)) {
            return item
        }
    }
    fmt.Printf(" '%s' not found!\n", title)
    return nil
}

func (l *Library) ShowAvailable() {
    fmt.Printf("\n Available Items in %s:\n", l.Name)
    for _, item := range l.items {
        if item.IsAvailable() {
            fmt.Printf("   %s\n", item.ItemType())
        }
    }
}

func PROJECTlM() {
    lib := Library{Name: "City Central Library"}

   
    B1 := &Book{baseitem: BaseItem{Title: "Python Crash Course",
        Author: "Eric Matthes", ItemID: "B001", available: true}, Pages: 560}
    B2 := &Book{BASEITEM: BaseItem{Title: "Clean Code",
        Author: "Robert Martin", ItemID: "B002", available: true}, Pages: 431}
    m1 := &Magazine{BASEITEM: BaseItem{Title: "Tech Today",
        Author: "John Doe", ItemID: "M001", available: true}, Issue: "March 2026"}
    d1 := &DVD{BaseItem: BASEITEM{Title: "The Social Network",
        Author: "David Fincher", ItemID: "D001", available: true}, Duration: 120}

    lib.AddItem(B1)
    lib.AddItem(B2)
    lib.AddItem(m1)
    lib.AddItem(d1)

  
    lib.ShowAll()

    B1.BORROW()
    B2.BORROW()     

    lib.ShowAvailable()

  
    B1.RETURNITEM()
    lib.ShowAvailable()
}