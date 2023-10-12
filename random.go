package main

import (
	"fmt"
	"math/rand"
)

func main() {

	randomNumber := rand.Intn(10) + 1
	playGuessingGame(randomNumber)
}

func playGuessingGame(randomNumber int) {
	fmt.Println("Welcome to the Game!")
	fmt.Println("Try to guess it.")

	for {
		var guess int
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 10.")
			continue
		}

		if guess < 1 || guess > 10 {
			fmt.Println("Please enter a number between 1 and 10.")
			continue
		}

		if guess == randomNumber {
			fmt.Println("It's Correct!", randomNumber)
			break
		} else if guess < randomNumber {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}
	}
}
