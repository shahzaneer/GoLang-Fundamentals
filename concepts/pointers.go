// Package concepts demonstrates basic Go programming concepts.
// This file explains the use of pointers in Go, including pass by value and pass by reference.

package concepts

import "fmt"

// Function demonstrating pass by value
// This function will take a copy of the variable and changes to it will not reflect outside the function.
func passByValue(x int) {
    fmt.Println("Inside passByValue function - Before modification:", x)
    x = 10 // Changing the value of x
    fmt.Println("Inside passByValue function - After modification:", x)
}

// Function demonstrating pass by reference
// This function will modify the original variable passed to it because we are passing a pointer.
func passByReference(x *int) {
    fmt.Println("Inside passByReference function - Before modification:", *x)
    *x = 20 // Dereferencing the pointer to modify the original value
    fmt.Println("Inside passByReference function - After modification:", *x)
}

// Function to demonstrate pointers
func pointerExample() {
    // Declaring a variable
    var a int = 5
    fmt.Println("Original value of a:", a)

    // Passing by value: a copy of a is passed to the function, so the changes inside the function won't affect the original.
    passByValue(a)
    fmt.Println("Value of a after passByValue:", a) // a is still 5

    // Passing by reference: passing the memory address of a, so the changes inside the function affect the original variable.
    passByReference(&a)
    fmt.Println("Value of a after passByReference:", a) // a is now 20
}

// Demonstrating the use of pointers
func pointerUsage() {
    // Declaring a pointer and initializing it
    var ptr *int
    fmt.Println("Value of pointer ptr before initialization:", ptr)

    // Allocating memory for the pointer
    num := 30
    ptr = &num
    fmt.Println("Pointer ptr points to:", ptr)   // Address of num
    fmt.Println("Value of num through pointer:", *ptr) // Dereferencing to get the value

    // Modifying the value through the pointer
    *ptr = 40
    fmt.Println("New value of num after dereferencing pointer:", num)
}

// Pointer with arrays
func pointerWithArrays() {
    arr := [3]int{10, 20, 30}
    fmt.Println("Original array:", arr)

    // Passing the pointer to the array
    modifyArray(&arr)
    fmt.Println("Modified array:", arr)
}

// Function that modifies an array via pointer
func modifyArray(arr *[3]int) {
    arr[0] = 100 // Modifying the array's first element via the pointer
}

func pointerMain() {
    // Demonstrating pointer basics
    pointerUsage()

    // Demonstrating pass by value and reference
    pointerExample()

    // Demonstrating pointer with arrays
    pointerWithArrays()
}
