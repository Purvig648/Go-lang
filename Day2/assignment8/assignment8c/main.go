package main

import (
	"projectstruct/model"
	"projectstruct/person"
)

func main() {
	person.PrintPerson(model.Person{
		Name: "Purvi",
		Age:  21,
	})
	// log.New()
}
