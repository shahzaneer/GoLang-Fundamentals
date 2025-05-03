// Package concepts demonstrates basic Go programming concepts.
// This file focuses solely on data types in Go.

package concepts

import "fmt"

// Data Types
// Go supports several basic data types like string, int, float64, bool, etc.

func dataTypes() {
    var integer int = 42               // Integer type
    var floatingPoint float64 = 3.1415  // Floating-point type
    var isAvailable bool = true         // Boolean type
    var greeting string = "Hello, Go!"  // String type

    // Output
    fmt.Println("Integer:", integer)
    fmt.Println("Floating Point:", floatingPoint)
    fmt.Println("Boolean:", isAvailable)
    fmt.Println("String:", greeting)
}
