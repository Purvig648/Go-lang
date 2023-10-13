package main

import (
	"assignments/models"
	"assignments/shape"
	"fmt"
)

func main() {
	fmt.Println("Area of circle is", shape.Area(models.Circle{Radius: 21}))

}
