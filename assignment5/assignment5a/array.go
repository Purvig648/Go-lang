package main

import "fmt"

func main() {
	var number [5]int
	number[0] = 21
	number[1] = 20
	number[2] = 15
	number[3] = 45
	number[4] = 23
	//fmt.Println(number)

	var num = [5]int{21, 22, 23, 24, 25}
	//fmt.Println("The numbers are", num)
	fmt.Println("The Numbers are")
	for d := 0; d < len(num); d++ {
		fmt.Println(num[d])
	}

}
