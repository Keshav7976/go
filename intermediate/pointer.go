package inte

import "fmt"

func main() {
	var ptr *int
	num:=10
	ptr = &num // Assigning the address of num to ptr
	// fmt.Println("Value of num:", num)         // Output: Value of num: 10
	// fmt.Println("Address of num:", &num)       // Output: Address of num: 0x...
	// fmt.Println("Value of ptr:", ptr)          // Output: Value of ptr: 0x...
	// fmt.Println("Value pointed by ptr:", *ptr) // Output: Value pointed by ptr: 10
	point(ptr) // Incrementing the value pointed by ptr
	fmt.Println("Value of num after increment:", num) // Output: Value of num after increment: 11
}
func point(ptr *int) {
	*ptr++
}