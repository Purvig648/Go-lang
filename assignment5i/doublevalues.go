package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for i := 0; i < len(a); i++ {
		a[i] = a[i] * 2

	}
	fmt.Println(a)

}
