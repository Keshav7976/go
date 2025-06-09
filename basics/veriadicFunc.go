package basics

import "fmt"

func main() {
	result:=sum(1, 2, 3, 4, 5)
	fmt.Println("Sum of numbers:", result)
	string1, total := sums("Total is:", 1, 2, 3, 4, 5)
	fmt.Println(string1, total)
	sequence,total1:=slicess(1, 10,20,30,40,50)
	fmt.Println("sequence:", sequence, "Has total:", total1)
	// veriadic functions to access slice values
	slice := []int{1, 2, 3, 4, 5}
	sequence2, total2 := slicess(2, slice...)
	fmt.Println("sequence2:", sequence2, "Has total2:", total2)
}

// variadic function example
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func sums(s string ,num...int) (string,int) {
	total := 0
	for _, number := range num {
		total += number
	}
	return s,total
}

func slicess(sequence int ,num...int) (int,int) {
	total := 0
	for _, number := range num {
		total += number
	}
	return sequence,total
}