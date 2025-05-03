// Package concepts demonstrates basic Go programming concepts.
// This file explains how to define and use structs in Go.

package concepts

import "fmt"

// Struct in Go
// A struct is a composite data type that groups together variables (fields) under one name. 
// These variables can be of different types, making structs very useful for modeling real-world entities.

// Defining a struct
type Person struct {
    Name    string
    Age     int
    Address string
}

// Function that takes a struct as an argument
func introduce(p Person) {
    fmt.Printf("Hello, my name is %s. I am %d years old, and I live at %s.\n", p.Name, p.Age, p.Address)
}

// Struct with methods
type Rectangle struct {
    Width  float64
    Height float64
}

// Method to calculate the area of the Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method to calculate the perimeter of the Rectangle
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Pointer receiver example
func (r *Rectangle) SetDimensions(width, height float64) {
    r.Width = width
    r.Height = height
}

// Anonymous structs
func anonymousStructExample() {
    // Anonymous struct is a struct without a named type
    person := struct {
        Name    string
        Age     int
        Address string
    }{
        Name:    "Alice",
        Age:     25,
        Address: "123 Main St.",
    }

    fmt.Printf("Anonymous Struct: %s, %d, %s\n", person.Name, person.Age, person.Address)
}

// Struct with nested structs
type Address struct {
    Street string
    City   string
    State  string
    Zip    string
}

type Employee struct {
    ID      int
    Name    string
    Address Address // Nested struct
}

func nestedStructExample() {
    emp := Employee{
        ID:   1,
        Name: "John Doe",
        Address: Address{
            Street: "456 Elm St.",
            City:   "Springfield",
            State:  "IL",
            Zip:    "62701",
        },
    }

    fmt.Printf("Employee: %d, %s, Address: %s, %s, %s, %s\n",
        emp.ID, emp.Name, emp.Address.Street, emp.Address.City, emp.Address.State, emp.Address.Zip)
}

func structExample() {
    // Example 1: Using a struct with fields
    fmt.Println("Struct Example 1:")
    person := Person{Name: "John", Age: 30, Address: "123 Main St."}
    introduce(person)

    // Example 2: Struct with methods
    fmt.Println("\nStruct Example 2:")
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Printf("Area: %f, Perimeter: %f\n", rect.Area(), rect.Perimeter())

    // Example 3: Pointer receiver
    fmt.Println("\nStruct Example 3:")
    rect.SetDimensions(20, 10)
    fmt.Printf("Updated dimensions: Width: %f, Height: %f\n", rect.Width, rect.Height)

    // Example 4: Anonymous struct
    fmt.Println("\nAnonymous Struct Example:")
    anonymousStructExample()

    // Example 5: Nested structs
    fmt.Println("\nNested Struct Example:")
    nestedStructExample()
}
