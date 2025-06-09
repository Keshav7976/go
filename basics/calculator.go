package basics

import (
	"fmt"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return a / b
}

func main() {
	var a, b float64
	var op string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&op)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	switch op {
	case "+":
		fmt.Printf("Result: %.2f\n", add(a, b))
	case "-":
		fmt.Printf("Result: %.2f\n", subtract(a, b))
	case "*":
		fmt.Printf("Result: %.2f\n", multiply(a, b))
	case "/":
		fmt.Printf("Result: %.2f\n", divide(a, b))
	default:
		fmt.Println("Invalid operator")
	}
}