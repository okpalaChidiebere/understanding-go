package main

import "fmt"

// shape interface
type Shape interface {
	getArea() float64
}

// type square implementing shape interface
type Square struct {
	sideLength float64
}

// type triangle implementing shape interface
type Triangle struct {
	height float64
	base   float64
}

func main() {

	var sq Square
	sq.sideLength = 12
	printArea(sq)

	tr := Triangle{base: 10, height: 10} //remember in go we like namong conventions for variables names to be the first 3 characters at most not "triangleShape"
	printArea(tr)

	printArea(sq)
	printArea(tr) //dont fall into the trap thinking you can perform triangleShape.printArea() :)

}

// Method getArea() of interface shape implemented by Square
func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

// Method getArea() of interface shape implemented by Triangle
func (t Triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func printArea(s Shape) {
	fmt.Printf("The Area for your shape is %v\n", s.getArea())
}
