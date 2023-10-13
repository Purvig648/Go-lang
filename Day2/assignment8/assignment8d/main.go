package main

import (
	"assignments/assignment8d/models"
	"assignments/assignment8d/shape"
	"fmt"
)

func main() {
	fmt.Println("Area of circle is", shape.Area(models.Circle{Radius: 21}))

}
