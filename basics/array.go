package basics

import "fmt"

func main() {
	number:=[5]int{1, 2, 3, 4, 5}
	var fruit [4]string
	fmt.Println(number)
	fruit[0] = "Apple"
	fruit[1] = "Banana"
	fruit[2] = "Cherry"
	fruit[3] = "Date"
	fmt.Println("fruits array:", fruit)
	for i:=0;i<len(number);i++ {
		fmt.Println("number at index", i, "is", number[i])
	}
	for index, value:=range number{
		fmt.Printf("Index:%d Value:%d\n", index, value)
	}
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{4, 5, 6}
	fmt.Println("array1 is equal to array2:", array1==array2) // false, arrays are not equal
	array3 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	fmt.Println("array3:", array3) // multi-dimensional array
	originalarray := [5]int{1, 2, 3, 4, 5}
	copiedarray := originalarray // copying an array
	copiedarray[0] = 10 // modifying copied array does not affect original
	fmt.Println("original array:", originalarray)
	fmt.Println("copied array:",copiedarray) // copied array is a separate copy
	// using pointers with arrays
	pointerArray := &originalarray // pointer to the original array
	pointerArray[0] = 20 // modifying through pointer affects original
	fmt.Println("Modified original array through pointer:", originalarray)
	fmt.Println("Pointer to original array:", pointerArray)
}