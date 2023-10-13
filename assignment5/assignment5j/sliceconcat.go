package main

import "fmt"

func main() {
	slice1 := []int{10, 20, 30, 40}
	slice2 := []int{50, 60, 70, 80}
	cap := len(slice1) + len(slice2)
	slice := make([]int, 0, cap)

	slice = append(slice, slice2...)
	fmt.Println(slice)
	slice = append(slice, slice1...)
	// slice = appennd([]slice2)
	// slice = append([]slice1)
	fmt.Println(slice)

}
