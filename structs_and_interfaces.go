// A struct is a type which contains named fields holding complex data
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

type Rect struct {
	Width, Height float64
	Name          string
}

type Cuboid struct {
	Rect   // Embedded type
	Length float64
}

func Area(c *Circle) float64 {
	// Area is a standard function
	// I make use of a pointer here instead of copying 'c'
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Area() float64 {
	// Area now is a method -- it allows us to call the function using the . operator
	return math.Pi * c.Radius * c.Radius
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (q Cuboid) Area() float64 {
	return 2 * (q.Width*q.Height + q.Width*q.Length +
		q.Height*q.Length)
}

type Shape interface {
	// Instead of defining fields, we define a “method set”.
	// A method set is a list of methods that a type must have in order to “implement” the interface.
	// Every struct which has an Area() method and returns int64
	// is considered as a 'Shape' object
	Area() float64
}

func measurePrint(s Shape) {
	fmt.Printf("%T, %+v\n", s, s)
	fmt.Printf("Area: %v\n\n", s.Area()) // Call method Area() on Shape
}

func main() {
	c := Circle{Radius: 5}
	r := Rect{Width: 2, Height: 3}
	q := Cuboid{Rect: r, Length: 5}
	fmt.Println(Area(&c)) // call standard function through pointer

	fmt.Println(c.Area()) // call the method Area associated with Rect
	fmt.Println(r.Area())
	fmt.Println(q.Area())

	// alternatively, make use of interface
	measurePrint(c) // pass "c" as a valid Shape
	measurePrint(r)
	measurePrint(q)

}
