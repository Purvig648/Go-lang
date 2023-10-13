package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}
type Circle struct {
	Radius float32
}
type Rectangle struct {
	Length  float32
	Breadth float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}
func (r Rectangle) Area() float32 {
	return r.Length * r.Breadth
}
func perform(s Shape) {
	fmt.Println("The area is ", s.Area())

}
func main() {
	c := Circle{
		Radius: 3.0,
	}
	r := Rectangle{
		Length:  6.0,
		Breadth: 5.0,
	}
	perform(c)
	perform(r)

}
