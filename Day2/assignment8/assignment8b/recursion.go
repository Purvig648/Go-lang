package main

import "fmt"

func fact(n int) int {
	//var s int = 1
	/*for i := 1; i < n; i++ {
		s = s + s*i
	}*/
	if n <= 1 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	var n int = 4
	fmt.Println(fact(n))
}
