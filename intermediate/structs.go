package intermidiate

import (
	"fmt"
	
)

func main() {
	
	p1 := person{name: "Alice",
	surn:"dune",
	 age: 30,
	 address:Address{
		street:"123 Main St",
		 city:"Wonderland",
		},
	homecell:homecell{
		home: "Wonder Home",
		cell: "123-456-7890",
	}}
	p2 := person{name: "Bob",surn:"theBuilder", age: 25}

	p2.address.street = "456 Elm St"
	p2.address.city = "Builderland"
	fmt.Println("Person 1:", p1.address.street, p1.address.city)
	fmt.Println("Person 1 Home:", p1.home)// Accessing embedded struct field directly without Address
	fmt.Println("Person 2:", p2.address.street, p2.address.city)
	user:=struct{
		name string
		mail string
	}{
		name: "John Doe",
		mail: "hbsinfiv@jnvvnfjdvn.org",
	}
	fmt.Println("User Mail:", user.mail)
	fmt.Println("User Name:", p2.fullName())
	p1.addage()
	fmt.Println("changing age of p1", p1.age)	
}
type person struct {
		name string
		surn string
		age  int
		address Address
		// address Address // embedded struct
		homecell
	}
type Address struct {
	street string
	city   string
}
type homecell struct {
	home string
	cell string
}	

func (p person) fullName() string{
	return p.name + " " + p.surn
}
func (p*person) addage(){
	p.age++
}