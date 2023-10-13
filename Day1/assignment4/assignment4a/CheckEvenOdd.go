package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter a number")
	fmt.Scan(&n)
	fmt.Println(CheckEvenOdd(n))

}
func CheckEvenOdd(n int) string {
	if n%2 == 0 {
		return "It is a Even number"
	}
	return "It is a odd number"
}
