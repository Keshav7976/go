package basics

import (
	"fmt"
)

func main() {
	// var numbers []int // Declare a slice of integers
	// numbers = []int{1, 2, 3, 4, 5} // Initialize the slice with values
	array := [5]int{1, 2, 3, 4, 5} 
	slice := array[1:4] // Create a slice from the array (elements 2, 3, 4)
	fmt.Println("Slice from array:", slice) 
	slice1:=append(slice, 6, 7) // Append elements to the slice
	fmt.Println("Slice after appending:", slice1)
	sliceCopy:=make([]int,len(slice1)) // Create a new slice with the same length as slice1
	copy(sliceCopy, slice1) // Copy elements from slice1 to sliceCopy
	fmt.Println("Copy of slice:", sliceCopy) // Print the copied slice
	// if slices.Equal[slice1,sliceCopy](
	// 	fmt.println("slice1 is equal to sliceCopy", slice1 == sliceCopy) // Check if the two slices are equal
	// )
	twod := make ([][]int, 3) // Create a two-dimensional slice
	for i := 0; i < 3; i++ {
		x:= i + 1 // Example value for each row
		twod[i] = make([]int, x) // Initialize each row with a slice of integers
		for j := 0; j < x; j++ {
			twod[i][j] = i + j // Fill the two-dimensional slice with values
		}
	}
	fmt.Println("Two-dimensional slice:", twod) // Print the two-dimensional slice

}