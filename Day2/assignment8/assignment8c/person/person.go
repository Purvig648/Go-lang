package person

import (
	"fmt"
	"projectstruct/model"
)

func PrintPerson(u model.Person) {
	fmt.Println("The name is ", u.Name)
	fmt.Println("The age is ", u.Age)

}
