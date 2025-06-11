package intermide

import "fmt"

type Rectangle struct {
	length int
	width  int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}
func (r *Rectangle) scale(factor int) {
	r.length *= factor
	r.width *= factor
}
func main() {
	rect := Rectangle{length: 10, width: 5}
	area := rect.Area()
	fmt.Println("Area of rectangle:", area)
	rect.scale(2)
	area = rect.Area()
	fmt.Println("Area of scaled rectangle:", area)
	num1 := Myint(10)
	num2 := Myint(-5)
	fmt.Println("Is num1 positive?", num1.isPositive())
	fmt.Println("Is num2 positive?", num2.isPositive())
}
type Myint int
func (m Myint) isPositive() bool {
	return m > 0
}