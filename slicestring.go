package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "orange", "strawberry", "mango"}
	fmt.Println(fruits)
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}
}
