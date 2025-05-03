// Package concepts demonstrates the use of conditional statements in Go.
package concepts

import "fmt"

func IfElseExamples() {
    fmt.Println("1. Basic if:")
    age := 17
    if age < 18 {
        fmt.Println("You're a minor.")
    }

    fmt.Println("\n2. if with else:")
    isRaining := false
    if isRaining {
        fmt.Println("Bring an umbrella.")
    } else {
        fmt.Println("Enjoy the sunshine!")
    }

    fmt.Println("\n3. if, else if, else:")
    temp := 35
    if temp < 0 {
        fmt.Println("Freezing cold!")
    } else if temp >= 0 && temp <= 30 {
        fmt.Println("Pleasant weather.")
    } else {
        fmt.Println("Too hot!")
    }

    fmt.Println("\n4. if with short statement:")
    if marks := 85; marks >= 90 {
        fmt.Println("Grade: A")
    } else if marks >= 75 {
        fmt.Println("Grade: B")
    } else {
        fmt.Println("Grade: C")
    }

    fmt.Println("\n5. Nested if:")
    age = 20
    hasID := true
    if age >= 18 {
        if hasID {
            fmt.Println("Access granted.")
        } else {
            fmt.Println("No ID, no entry.")
        }
    } else {
        fmt.Println("Too young.")
    }

    fmt.Println("\n6. Logical operators:")
    age = 25
    citizen := true
    if age >= 18 && citizen {
        fmt.Println("You can vote.")
    }
    if age < 18 || !citizen {
        fmt.Println("You cannot vote.")
    }
}
