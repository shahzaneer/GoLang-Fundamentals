// Package concepts demonstrates basic Go programming concepts.
// This file explains all types of loops in Go.

package concepts

import "fmt"

// Loops in Go
// Go only has one looping construct, the `for` loop, which can be used in different ways
// to achieve various loop types (traditional, while-like, and infinite).

func loops() {
    // **1. Traditional for loop (like in C, Java, etc.)**
    fmt.Println("Traditional For Loop Example:")

    // A simple for loop that runs from 1 to 5
    for i := 1; i <= 5; i++ {
        fmt.Println(i) // Output: 1, 2, 3, 4, 5
    }

    // **2. While-like for loop (condition only)**
    fmt.Println("\nWhile-like For Loop Example:")

    // A `for` loop that mimics a while loop (no initialization or increment)
    j := 1
    for j <= 5 {
        fmt.Println(j) // Output: 1, 2, 3, 4, 5
        j++
    }

    // **3. Infinite For Loop**
    fmt.Println("\nInfinite For Loop Example:")

    // A `for` loop without a condition creates an infinite loop.
    // Be careful with this one, as it will run forever unless you explicitly break the loop.
    // Uncomment the following lines to test it out.

    // k := 1
    // for {
    //     fmt.Println(k)
    //     k++
    //     if k > 5 {
    //         break // Exit the infinite loop when k exceeds 5
    //     }
    // }

    // **4. For loop with `range` (iterating over collections)**
    fmt.Println("\nFor Loop with Range Example:")

    // `range` is used to iterate over arrays, slices, strings, maps, and channels.
    nums := []int{10, 20, 30, 40, 50}
    for index, value := range nums {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    // Example of iterating over a string using `range`
    str := "GoLang"
    fmt.Println("\nIterating over string with range:")
    for index, char := range str {
        fmt.Printf("Index: %d, Character: %c\n", index, char)
    }

    // **5. For loop with `continue` and `break`**
    fmt.Println("\nFor Loop with Continue and Break Example:")

    // `continue` skips the current iteration and continues to the next iteration.
    // `break` breaks the loop entirely.

    for i := 1; i <= 10; i++ {
        if i == 5 {
            continue // Skips printing 5
        }
        if i == 8 {
            break // Stops the loop at 8
        }
        fmt.Println(i) // Output: 1, 2, 3, 4, 6, 7
    }
}
