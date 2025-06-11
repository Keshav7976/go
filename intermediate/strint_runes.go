package intermidiate

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message:="Hello,\nGo!" 
	message1:="Hello,\tGo!" 
	message2:="Hello,Go!"
	message3:=`Hello,\nGo!`
	fmt.Println(message)   
	fmt.Println(message1)  
	fmt.Println(message2)  
	fmt.Println(message3)  
	fmt.Println("Length of message:", len(message))
	fmt.Println("Length of message1:", len(message2))
	greet:="Hello"
	name:="alice"
	fmt.Println(greet + " " + name) // Concatenation
	str1:="Apple" // ASCI values: A=65, p=112, l=108, e=101
	str2:="banana"// ASCI values: b=98, a=97, n=110
	str3:="app"// ASCI values: a=97, p=112, p=112

	// compares on the basis of ASCII values
	fmt.Println(str1< str2) 
	fmt.Println(str3> str1)
	fmt.Println(str3> str2)

	// for i , char:=range message{
	// 	fmt.Printf("character at index %d is %c\n",i,char)
	// 	fmt.Printf("%v\n", char) // Print the rune value
	// }
	fmt.Println("Length of message in runes:", len([]rune(message)))
	fmt.Println("Rune count:",utf8.RuneCountInString(message1))
	var ch rune='a'
	fmt.Println(ch)
	cstr:=string(ch) // Convert rune to string
	fmt.Println(cstr) // Output: a

	fmt.Printf("Type of rune: %T\n", ch) 
	fmt.Printf("Type of cstr: %T\n", cstr)

	jhello:="こんにちは" // Japanese for "Hello"
	fmt.Println("Japanese Hello:", jhello)
	for _,jrune:=range jhello {
		fmt.Printf("Character: %c\n", jrune)
	}

}