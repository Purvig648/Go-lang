package main

import "fmt"

func main() {
	fruits := make(map[string]int)
	fruits["Apple"] = 20
	fruits["Banana"] = 10
	fruits["Orange"] = 15
	fruits["Grapes"] = 25
	for k, v := range fruits {
		fmt.Println(k, v)
	}
}
