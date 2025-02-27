package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(`
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)
	`)

	fmt.Print("Enter your choice (1, 2, or 3): ")
	var choice uint8
	fmt.Scan(&choice)

	switch choice {
	case 1:
		startGame("Easy", 10)
	case 2:
		startGame("Medium", 5)
	case 3:
		startGame("Hard", 3)
	default:
		fmt.Println("Unknown difficulty level. Exiting game.")
	}
}

func startGame(difficulty string, chances uint8) {
	fmt.Printf("\nGreat! You selected %s difficulty. You have %d chances.\n\n", difficulty, chances)

	targetNumber := generateRandomNumber()

	for i := uint8(0); i < chances; i++ {
		fmt.Printf("Attempt %d/%d - Enter your guess: ", i+1, chances)
		var guess int
		fmt.Scan(&guess)

		if guess < 1 || guess > 100 {
			fmt.Println("Please enter a number between 1 and 100.")
			i-- // Don't count invalid attempts
			continue
		}

		if guess < targetNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > targetNumber {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("🎉 Congratulations! You guessed the correct number!")
			return
		}
	}

	fmt.Printf("\n😢 You've run out of attempts! The correct number was %d.\n", targetNumber)
}

func generateRandomNumber() int {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	return rng.Intn(100) + 1
}
