package main

import "fmt"

func main() {
	studentGrades := make(map[string]float64)
	studentGrades["Purvi"] = 8.5
	studentGrades["Jeevan"] = 9
	studentGrades["Aftab"] = 9.5
	studentGrades["Ashwini"] = 8.0
	for k, v := range studentGrades {
		fmt.Println(k, v)
	}
	delete(studentGrades, "Ashwini")
	fmt.Println("The elements after deleting")
	for k, v := range studentGrades {
		fmt.Println(k, v)
	}

}
