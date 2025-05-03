// Package concepts demonstrates basic Go programming concepts.
// This file explains the use of interfaces in Go.

package concepts

import "fmt"

// Interface in Go
// An interface is a type that specifies a set of method signatures (behavior) but does not provide implementations.
// Any type that implements these methods implicitly satisfies the interface. This is called "duck typing."

// Defining an interface
type Speaker interface {
    Speak() string
}

// Defining a type that implements the interface
type PersonD struct {
    Name string
}

// Implementing the Speak method for the Person type
func (p PersonD) Speak() string {
    return fmt.Sprintf("Hello, my name is %s.", p.Name)
}

// Another type implementing the same interface
type Animal struct {
    Species string
}

// Implementing the Speak method for the Animal type
func (a Animal) Speak() string {
    return fmt.Sprintf("I am a %s.", a.Species)
}

// Function that accepts any type that satisfies the Speaker interface
func introduceR(speaker Speaker) {
    fmt.Println(speaker.Speak())
}

// Empty interface example
func emptyInterfaceExample() {
    var anything interface{} = 42
    fmt.Println(anything) // Can hold any type of value (string, int, etc.)
}

// Type assertion example
func typeAssertionExample() {
    var anyValue interface{} = 42
    
    // Performing type assertion
    if value, ok := anyValue.(int); ok {
        fmt.Printf("Value is an int: %d\n", value)
    } else {
        fmt.Println("Value is not an int.")
    }
}

// Type switch example
func typeSwitchExample() {
    var x interface{} = "Hello, Go!"
    
    switch v := x.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    default:
        fmt.Println("Unknown type")
    }
}

func interfaceExample() {
    // Example 1: Using the interface
    fmt.Println("Interface Example 1:")
    person := Person{Name: "John"}
    animal := Animal{Species: "Lion"}
    
    // Passing types that implement the Speaker interface
    introduce(person)
    introduceR(animal)
    
    // Example 2: Empty interface
    fmt.Println("\nEmpty Interface Example:")
    emptyInterfaceExample()
    
    // Example 3: Type assertion
    fmt.Println("\nType Assertion Example:")
    typeAssertionExample()
    
    // Example 4: Type switch
    fmt.Println("\nType Switch Example:")
    typeSwitchExample()
}
