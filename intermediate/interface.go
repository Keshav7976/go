package intermidiate
import "fmt"
type geometry interface{
	area() float64
	peri() float64
}
type rect struct {
	length, width float64
}
type circle struct {
	radius float64
}
func (r rect) area() float64 {
	return r.length * r.width
}
func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func (r rect) peri() float64 {
	return 2 * (r.length + r.width)
}
func (c circle) peri() float64 {
	return 2 * 3.14 * c.radius
}
func measure(g geometry) {
	fmt.Println("Geometry Type:", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.peri())
}
func main(){
	r:=rect{length: 10, width: 5}
	c := circle{radius: 7}
	measure(r)
	measure(c)
}