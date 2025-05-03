// Package concepts demonstrates basic Go programming concepts.
// This file explains the use of switch statements in Go.

package concepts

import "fmt"

// Switch statement in Go
// The switch statement is an alternative to multiple if-else chains. 
// It simplifies the syntax and makes the code more readable.

func switchStatement() {
    // **1. Basic switch example**
    fmt.Println("Basic Switch Example:")

    day := 3
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    case 4:
        fmt.Println("Thursday")
    case 5:
        fmt.Println("Friday")
    case 6:
        fmt.Println("Saturday")
    case 7:
        fmt.Println("Sunday")
    default:
        fmt.Println("Invalid day")
    }

    // **2. Switch with multiple cases for a single statement**
    fmt.Println("\nSwitch with Multiple Cases Example:")

    fruit := "banana"
    switch fruit {
    case "apple", "banana", "cherry":
        fmt.Println("It's a fruit!")
    default:
        fmt.Println("Unknown fruit")
    }

    // **3. Switch with no expression (acts like if-else)**
    fmt.Println("\nSwitch with No Expression Example:")

    num := 10
    switch {
    case num%2 == 0:
        fmt.Println("Even number")
    case num%2 != 0:
        fmt.Println("Odd number")
    }

    // **4. Fallthrough in switch**
    fmt.Println("\nFallthrough Switch Example:")

    score := 90
    switch {
    case score >= 90:
        fmt.Println("Excellent!")
        fallthrough // Fallthrough will execute the next case as well
    case score >= 75:
        fmt.Println("Good job!")
    case score >= 50:
        fmt.Println("Keep trying!")
    default:
        fmt.Println("Needs improvement")
    }

    // **5. Type switch (to switch on type, not value)**
    fmt.Println("\nType Switch Example:")

    var item interface{} = 42
    switch v := item.(type) {
    case int:
        fmt.Println("Item is an integer:", v)
    case string:
        fmt.Println("Item is a string:", v)
    case bool:
        fmt.Println("Item is a boolean:", v)
    default:
        fmt.Println("Unknown type")
    }

    // **6. Switch with expressions (case conditions)**
    fmt.Println("\nSwitch with Case Conditions Example:")

    num2 := 20
    switch {
    case num2 < 10:
        fmt.Println("Number is less than 10")
    case num2 >= 10 && num2 <= 20:
        fmt.Println("Number is between 10 and 20")
    case num2 > 20:
        fmt.Println("Number is greater than 20")
    }
}
