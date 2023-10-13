package main

import "fmt"

func main() {
	var sum int = 0
	numbers := []int{10, 20, 32, 12, 7, 8, 9, 3}
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			sum = sum + numbers[i]
		}
	}
	fmt.Println("The sum of even numbers is: ", sum)
}
