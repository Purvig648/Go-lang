package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	Age  int
	City string
}

func main() {
	E1 := Employee{01, "Purvi", 21, "Bangalore"}
	E2 := Employee{02, "Jeevan", 22, "Chitradurga"}
	E3 := Employee{013, "Aftab", 22, "Kushalnagar"}

	fmt.Println(E1)
	fmt.Println(E2)
	fmt.Println(E3)

}
