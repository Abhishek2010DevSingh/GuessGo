Here's your updated README with the attribution added:  

---

# Number Guessing Game

A simple number guessing game written in Go. The game randomly selects a number between 1 and 100, and the player has to guess it within a limited number of attempts based on the chosen difficulty level.

## Features
- Three difficulty levels:
  - **Easy**: 10 attempts
  - **Medium**: 5 attempts
  - **Hard**: 3 attempts
- Input validation to ensure guesses are within the correct range.
- Random number generation for each game session.
- User-friendly text-based interface.

## Installation
1. Ensure you have Go installed on your system. You can download it from [golang.org](https://golang.org/).
2. Clone this repository or create a new Go file and copy the code.
3. Navigate to the project directory.

## Usage
Run the game using the following command:

```sh
go run main.go
```

Follow the on-screen instructions to play the game.

## How to Play
1. Select a difficulty level by entering `1`, `2`, or `3`.
2. Enter your guess (a number between 1 and 100).
3. Receive feedback on whether your guess is too high or too low.
4. Continue guessing until you find the correct number or run out of attempts.
5. If you guess correctly, you win! Otherwise, the correct number is revealed at the end.

## Example Run
```
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)

Enter your choice (1, 2, or 3): 2

Great! You selected Medium difficulty. You have 5 chances.

Attempt 1/5 - Enter your guess: 50
Too low! Try again.

Attempt 2/5 - Enter your guess: 75
Too high! Try again.

Attempt 3/5 - Enter your guess: 60
Too low! Try again.

Attempt 4/5 - Enter your guess: 65
🎉 Congratulations! You guessed the correct number!
```

## License
This project is open-source and available under the MIT License.

## Contributions
Contributions are welcome! Feel free to submit pull requests or suggest improvements.

## Credits
This project idea was inspired by [roadmap.sh's Number Guessing Game](https://roadmap.sh/projects/number-guessing-game).

## Author
Created by Abhishek Kumar Singh

---

Let me know if you need any further changes! 🚀
