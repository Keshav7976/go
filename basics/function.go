package basics

import "fmt"

func main() {
	// fuc <name>(parametre list) returnType {
	// 	// function body
	// returnType
	// }
	/*greet :=func(){
		fmt.Println("this is a anonymous function")
	}
	greet()
	sum := add(5, 10)
	operation:=add
	result:=operation(12,33)
	fmt.Println("Result of add function:", result)
	fmt.Println("Sum:", sum)*/
	result := multipleoperations(5, 10, add)
	fmt.Println("Result of multiple operations:", result)
	multiplyBy2:=createMultiplier(2)
	fmt.Println("Multiply 5 by 2:", multiplyBy2(5))

}
func add(a int, b int) int {
	return a + b
}

// function that takes a function as a parameter
func multipleoperations(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}
// function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}