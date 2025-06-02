package basics
import "fmt"
func main() {
	var a,b int=10, 3
	var result int 
	result = a + b
	fmt.Println("Addition:", result)
	result = a - b
	fmt.Println("Subtraction:", result)
	result = a * b
	fmt.Println("Multiplication:", result)
	result = a / b
	fmt.Println("Division:", result)
	result = a % b
	fmt.Println("Modulus:", result)
	const pi float64 = 22.0 / 7.0
	fmt.Println("Value of Pi:", pi)
}