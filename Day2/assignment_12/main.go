package main

import "fmt"

type Animal interface {
	MakeSound() string
}
type Dog struct {
	DogSound string
}
type Cat struct {
	CatSound string
}

func (d Dog) MakeSound() string {
	return "Bow Bow"
}
func (c Cat) MakeSound() string {
	return "Meow Meow"
}
func Sound(A Animal) {
	fmt.Println("The sound is", A.MakeSound())
}
func main() {
	d := Dog{DogSound: "Lab"}
	c := Cat{CatSound: "perssian"}
	Sound(d)
	Sound(c)

}
