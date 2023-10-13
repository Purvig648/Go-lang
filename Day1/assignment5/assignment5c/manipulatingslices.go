package main

import (
	"fmt"
)

func main() {
	var s []int
	fmt.Println(s)
	s = append(s, 5)
	fmt.Println(s)
	s = append(s, 10)
	fmt.Println(s)
	s = append(s, 15)
	fmt.Println(s)
	s = append(s, 20)
	fmt.Println(s)
	s = append(s, 25)
	fmt.Println(s)
	fmt.Println("The length: ", len(s), "and capacity is: ", cap(s))
	var num int
	fmt.Println("The index to remove")
	fmt.Scan(&num)
	s = append(s[:num], s[num+1:]...)
	fmt.Println(s)

}
