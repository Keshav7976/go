package intermediate

import "fmt"

func main() {
	fact:=recursion(5)
	fmt.Println("Factorial of 5 is:", fact)
	fmt.Println("Sum of digits of 12345 is:", sumDigit(12345))
}
func recursion (n int) int{
	if n<=1{
		return n
	}else{
		return n * recursion(n-1)
	}

}
func sumDigit(n int) int{
	if n == 0 {
		return 0
	} else {
		return n%10 + sumDigit(n/10)
	}
}