package main

import (
	"assignments/assignment3/assignment3a/calculattor"
	"fmt"
)

func main() {
	fmt.Println("The square of numbers is: ", calculattor.Square(12))
	fmt.Println("The sum of numbers is: ", calculattor.Sum(12, 6))
	fmt.Println("The difference of numbers is: ", calculattor.Diff(12, 6))
	fmt.Println("The product of numbers is: ", calculattor.Prod(12, 6))

}
