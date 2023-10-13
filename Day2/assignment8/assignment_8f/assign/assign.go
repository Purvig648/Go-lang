package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func UpdateAge(s *Student) {
	s.Age = s.Age + 1
}
func main() {
	s := Student{
		Name: "Purvi",
		Age:  22,
	}
	fmt.Println(s)
	UpdateAge(&s)
	fmt.Println(s)

}
