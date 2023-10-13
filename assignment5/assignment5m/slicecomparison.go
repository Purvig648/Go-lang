package main

import "fmt"

func compareslice(slice1 []int, slice2 []int) bool {
	if len(slice2) != len(slice1) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
func main() {
	slice1 := []int{10, 20, 30, 40, 50}
	slice2 := []int{10, 20, 10, 40, 50}
	ok := compareslice(slice1, slice2)
	if !ok {
		fmt.Println("Slices are not equal")
		return
	}
	fmt.Println("Slices are equal")

}
