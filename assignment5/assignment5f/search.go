package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s = []int{10, 20, 30, 40, 50}
	//var searchele int
	searchele := os.Args[1]
	num, err := strconv.Atoi(searchele)
	if err != nil {
		fmt.Println("Invalid type")
		return
	}
	var found bool
	for i := 0; i < len(s); i++ {
		if s[i] == num {
			fmt.Println("Found it at index ", i)
			found = true
			break

		}
	}
	if !found == true {
		fmt.Println("Element not found", num)
	}
}
