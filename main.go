package main

import "fmt"

func main() {
	msg := `
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have 5 chances to guess the correct number.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)
	`

	fmt.Println(msg)

	fmt.Print("Enter your choice (1, 2, or 3):")
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
		fmt.Println("Unknown difficulty level.")
	}
}

func startGame(difficultyLevel string, chance uint8) {

}
