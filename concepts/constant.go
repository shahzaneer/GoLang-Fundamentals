// File: constant.go
package concepts

import "fmt"

func ShowConstantsConcept() {
	// -------------------------------------
	// 1. Basic constant declaration
	// Constants are immutable values known at compile time
	// -------------------------------------
	const Pi = 3.14
	const AppName = "Go Playground"
	fmt.Println("Pi:", Pi)
	fmt.Println("App Name:", AppName)

	// -------------------------------------
	// 2. Typed vs Untyped constants
	// Untyped constants take the type of the variable they are assigned to
	// -------------------------------------
	const untyped = 10      // Untyped constant
	const typed int = 100   // Typed constant

	var x float64 = untyped // allowed: 10 is untyped, gets converted
	// var y float64 = typed  // error: cannot use int constant in float context
	fmt.Println("Untyped assigned to float64:", x)

	// -------------------------------------
	// 3. Constant expressions
	// Constants can be used in expressions evaluated at compile time
	// -------------------------------------
	const a = 5
	const b = 10
	const sum = a + b
	fmt.Println("Sum of constants:", sum)

	// -------------------------------------
	// 4. iota — a Go constant generator
	// iota starts at 0 and increments by 1 automatically
	// Often used for enums or successive values
	// -------------------------------------
	const (
		First = iota // 0
		Second       // 1
		Third        // 2
	)
	fmt.Println("Iota values:", First, Second, Third)

	// -------------------------------------
	// 5. Skipping iota values
	// Use underscore _ to skip a value
	// -------------------------------------
	const (
		_    = iota             // skip 0
		Read = 1 << iota        // 1 << 1 = 2
		Write                  // 1 << 2 = 4
		Execute                // 1 << 3 = 8
	)
	fmt.Println("Permissions:", Read, Write, Execute)

	// -------------------------------------
	// 6. Constants with string iota pattern
	// Not automatically incremented like numbers, but can be used creatively
	// -------------------------------------
	const (
		North = "North"
		East  = "East"
		South = "South"
		West  = "West"
	)
	fmt.Println("Directions:", North, East, South, West)
}
