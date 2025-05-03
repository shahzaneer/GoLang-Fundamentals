// File: variables.go
package concepts

import "fmt"

func ShowVariablesConcept() {
	// -----------------------------
	// 1. Declaring a variable using var keyword
	// -----------------------------
	var age int = 25
	fmt.Println("Age:", age)

	// -----------------------------
	// 2. Type inference using var
	// -----------------------------
	var city = "Lahore"
	fmt.Println("City:", city)

	// -----------------------------
	// 3. Short variable declaration using :=
	// Only works inside functions
	// -----------------------------
	name := "Shahzaneer"
	fmt.Println("Name:", name)

	// -----------------------------
	// 4. Declaring multiple variables at once
	// -----------------------------
	var a, b, c int = 1, 2, 3
	fmt.Println("Multiple variables:", a, b, c)

	// -----------------------------
	// 5. Zero values in Go
	// Variables declared without a value get default zero values
	// -----------------------------
	var salary int
	var isActive bool
	var jobTitle string

	fmt.Println("Default int (salary):", salary)
	fmt.Println("Default bool (isActive):", isActive)
	fmt.Println("Default string (jobTitle):", jobTitle)

	// -----------------------------
	// 6. Constants
	// -----------------------------
	const Pi = 3.14
	const AppName = "GoLang Demo"
	fmt.Println("Constants:", Pi, AppName)

	// -----------------------------
	// 7. Block declaration (grouped variables)
	// -----------------------------
	var (
		username string = "admin"
		password string = "1234"
		port     int    = 8080
	)
	fmt.Println("Block declaration:", username, password, port)
}
