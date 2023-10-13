package main

import (
	"fmt"
)

func findindex(numbers []int, numb int) int {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == numb {
			return i
		}

	}
	return -1

}

func main() {
	numbers := []int{10, 20, 30, 40, 50, 60}
	fmt.Println("Enter the number you want to find")
	var numb int
	fmt.Scan(&numb)
	index := findindex(numbers, numb)
	fmt.Println("The", numb, "Found at index", index)

}
