package basics

import "fmt"

func main() {
	// for i := 0; i <5; i++{
	// 	fmt.Println(i)
	// }
	// numbers := []int{1, 2, 3, 4, 5}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }
	rows := 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {	
			fmt.Print(" ")
		}
		for k := 1; k <= i*2-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
