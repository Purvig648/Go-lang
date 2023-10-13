package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	evencount := 0
	oddcount := 0
	for _, num := range number {
		if num%2 == 0 {
			evencount++
		} else {
			oddcount++
		}
	}
	fmt.Println("The no of even numbers", evencount)
	fmt.Println("The no of odd numbers", oddcount)

}
