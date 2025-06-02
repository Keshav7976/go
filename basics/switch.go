package basic 

import "fmt"

func main() {
	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Start of the work week!")
	case "Friday":
		fmt.Println("Almost the weekend!")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Just another day.")
	}
	number:=2
	switch {
	case number>1:
		fmt.Println("Number is greater than 1")
		fallthrough // This will continue to the next case
	case number==2:
		fmt.Println("Number is equal to 2")
	default:
		fmt.Println("number is either greater than2 or lessa than 1")
	}
	checkType("hello")
	checkType(false)
}
func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("x is an int")
	case string:
		fmt.Println("x is a string")
		case float64:
		fmt.Println("x is a float64")
		default:
			fmt.Println("x is a unknown type")
		}
		
	}