package basics

import (
	"fmt"
	"maps"
)

func main() {
	mymap := make(map[string]int) // Create a map with string keys and int values
	mymap["apple"] = 5
	mymap["banana"] = 3
	mymap["cherry"] = 8
	fmt.Println("Map contents:", mymap) // Print the map contents

	delete(mymap, "banana") // Delete an entry from the map

	fmt.Println("Map after deletion:", mymap) // Print the map after deletion

	mymap["apple"] = 10 // Update an existing entry

	fmt.Println("Map after update:", mymap) // Print the map after update

	// clear(mymap) // Clear the map
	// fmt.Println("Map after clearing:", mymap) // Print the map after clearing

	value,ok :=mymap["apple"]
	fmt.Println("Value for 'apple':", value, "isOk:", ok) // Check if the key exists

	for k,v:=range mymap {
		fmt.Println(k , v) // Iterate over the map and print key-value pairs
	}

	if ok{
		fmt.Println("Key 'apple' exists in the map with value:", value) // Print if the key exists
	} else {
		fmt.Println("Key 'apple' does not exist in the map") // Print if the key does not exist
	}

	fmt.Println("Length of the map:", len(mymap)) // Print the length of the map

	map2 := map[string]int{"orange": 4, "grape": 6} 
	if maps.Equal(mymap, map2) {
		fmt.Println("Maps are equal") // Check if two maps are equal
	} else {
		fmt.Println("Maps are not equal") 
	}
	map3:=make(map[string]map[string]int) // Create a nested map
	map3["fruits"] = make(map[string]int)
	map3["fruits"]["kiwi"] = 12
	// map3["fruits"]= map2
	fmt.Println("Nested map:", map3) // Print the nested map
}