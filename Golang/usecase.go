package main

import "fmt"

// ABSTRACTION
type Employee interface {
    Role() string
    CalculateSalary() float64
    DisplayPayslip()
}

// ENCAPSULATION
type BaseEmployee struct {
    Name       string
    EmpID      string
    baseSalary float64   // private
}

func (b BaseEmployee) GetBaseSalary() float64 { return b.baseSalary }

// COMPOSITION
type FullTimeEmployee struct {
    BaseEmployee
    bonus float64
}

type PartTimeEmployee struct {
    BaseEmployee
    hoursWorked float64
}

type Contractor struct {
    BaseEmployee
    daysWorked float64
}

// FullTime methods
func (f FullTimeEmployee) Role() string { return "Full Time Employee" }
func (f FullTimeEmployee) CalculateSalary() float64 {
    return f.GetBaseSalary() + f.bonus
}
func (f FullTimeEmployee) DisplayPayslip() {
    fmt.Printf("\n %s (%s) | %s | ₹%.0f\n",
        f.Name, f.EmpID, f.Role(), f.CalculateSalary())
}

// PartTime methods
func (p PartTimeEmployee) Role() string { return "Part Time Employee" }
func (p PartTimeEmployee) CalculateSalary() float64 {
    return p.GetBaseSalary() * p.hoursWorked
}
func (p PartTimeEmployee) DisplayPayslip() {
    fmt.Printf("\n💼 %s (%s) | %s | ₹%.0f\n",
        p.Name, p.EmpID, p.Role(), p.CalculateSalary())
}

// Contractor methods
func (c Contractor) Role() string { return "Contractor" }
func (c Contractor) CalculateSalary() float64 {
    return c.GetBaseSalary() * c.daysWorked
}
func (c Contractor) DisplayPayslip() {
    fmt.Printf("\n💼 %s (%s) | %s | ₹%.0f\n",
        c.Name, c.EmpID, c.Role(), c.CalculateSalary())
}

func usecase() {
    // POLYMORPHISM
    employees := []Employee{
        FullTimeEmployee{BaseEmployee: BaseEmployee{Name: "Ravi", EmpID: "E001", baseSalary: 50000}, bonus: 10000},
        PartTimeEmployee{BaseEmployee: BaseEmployee{Name: "Sara", EmpID: "E002", baseSalary: 500}, hoursWorked: 80},
        Contractor{BaseEmployee: BaseEmployee{Name: "John", EmpID: "E003", baseSalary: 3000}, daysWorked: 20},
    }

    for _, e := range employees {
        e.DisplayPayslip()
    }
}