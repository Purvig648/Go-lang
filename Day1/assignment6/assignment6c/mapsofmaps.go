package main

import "fmt"

func main() {
	studentData := make(map[string]map[string]interface{})
	studentData["Purvi"] = map[string]interface{}{
		"Age":   21,
		"Grade": 9.2,
		"City":  "Bangalore",
	}
	studentData["Jeevan"] = map[string]interface{}{
		"Age":   22,
		"Grade": 9.5,
		"City":  "Chitradurga",
	}
	studentData["Aftab"] = map[string]interface{}{
		"Age":   22,
		"Grade": 9.7,
		"City":  "Kushalnagar",
	}
	for k, v := range studentData {
		fmt.Println(k, v)
	}

}
