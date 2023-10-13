package main

import "fmt"

type Address struct {
	Street  string
	City    string
	ZipCode int
}
type Person struct {
	Name string
	Address
}

func main() {
	P := Person{
		Name: "Purvi",
		Address: Address{
			Street:  "Aradhana Layout",
			City:    "Bangalore",
			ZipCode: 560076,
		},
	}
	fmt.Println(P)
}
