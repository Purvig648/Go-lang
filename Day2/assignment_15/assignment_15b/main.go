package main

import (
	"fmt"
	"time"
)

func sender(channel chan int) {
	for i := 1; i <= 5; i++ {
		channel <- i
	}
	close(channel)
}

func receiver(channel chan int) {
	for num := range channel {
		fmt.Printf("Received: %d\n", num)
	}
}

func main() {
	channel := make(chan int)

	go sender(channel)
	go receiver(channel)

	time.Sleep(2 * time.Second)
}
