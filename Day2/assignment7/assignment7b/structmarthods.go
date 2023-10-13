package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (u Rectangle) Area() int {
	area := u.Height * u.Width
	return area

}
func (u Rectangle) Perimetere() int {
	perimeter := 2 * (u.Height + u.Width)
	return perimeter

}

func main() {
	a := Rectangle{22, 21}
	fmt.Println("The area of a rectangle", a.Area())
	fmt.Println("The perimeter of a rectangle", a.Perimetere())
}
