package basics

import (
	"errors"
	"fmt"
)

func main() {
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)
	s, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(s)
	}
}
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}
func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if a < b {
		return "a is less than b", nil
	} else {
		return "", errors.New("a is equal to b")
	}
}
