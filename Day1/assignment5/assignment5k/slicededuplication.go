package main

import "fmt"

func remove(j int, numbers []int) []int {
	numbers = append(numbers[:j], numbers[j+1:]...)
	return numbers

}

func main() {
	numbers := []int{10, 10, 30, 40, 50, 20, 20}
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] == numbers[j] {
				numbers = remove(j, numbers)

			}

		}
	}
	fmt.Println(numbers)

}
