package main

import "fmt"

func main() {
	s := []int{10, 20, 40, 30, 28}
	var temp int
	s = append(s, 46)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				temp = s[i]
				s[i] = s[j]
				s[j] = temp
			}
		}
	}
	fmt.Println("The maximum element is", s[len(s)-1])
}
