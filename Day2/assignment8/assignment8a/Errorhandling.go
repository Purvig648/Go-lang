package main

import (
	"fmt"
	"os"
)

func main() {
	err := check("b.txt")
	if err != nil {
		fmt.Println("The file does not exist")
		fmt.Println(err)
		return
	}
	fmt.Println("The file is removed")

}
func check(file string) error {
	return os.Remove(file)
}
