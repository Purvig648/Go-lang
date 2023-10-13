package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}
func main() {

	nums := []int{10, 20, 30, 40}
	sum(nums...)
	sum(30, 40)
	sum(4, 5)

}
