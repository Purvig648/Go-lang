package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50}
	sum := 0
	for i := 0; i < len(s); i++ {
		sum = sum + s[i]
	}
	avg := (sum / (len(s)))
	fmt.Println("The sum is ", sum)
	fmt.Println("The average is: ", avg)

}
