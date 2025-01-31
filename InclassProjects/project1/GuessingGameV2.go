package main

import (
	"fmt"
	"math/rand/v2"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Welcome to the guessing game!")

	var Numtries int
	fmt.Println("Enter the maximum number of tries: ")
	fmt.Scan(&Numtries)
	var score int = Numtries
	var max int
	var goal int

	var difficulty int
	fmt.Println("Enter Wanted difficulty: 0 for Easy (1-50), 1 for Medium (1-100), 2 for Hard (1-200)")
	for {
		fmt.Scan(&difficulty)
		if difficulty < 0 || difficulty > 2 {
			fmt.Println("Invalid choice. Please enter 0, 1, or 2.")
		} else {
			fmt.Printf("Difficulty selected: %d\n", difficulty)
			break
		}
	}

	switch difficulty {
	case 0:
		goal = rand.IntN(50) + 1
		max = 50
	case 1:
		goal = rand.IntN(100) + 1
		max = 100
	case 2:
		goal = rand.IntN(200) + 1
		max = 200
	}

	for {
		var userguess int
		for i := 0; i < Numtries; i++ {
			fmt.Printf("Enter a number between 1 and %d: ", max)
			fmt.Scan(&userguess)

			if userguess < 1 || userguess > max {
				fmt.Printf("Invalid number! Enter a number between 1 and %d.\n", max)
				i--
				continue
			}

			if userguess < goal {
				fmt.Println("Too low! Try again.")
			} else if userguess > goal {
				fmt.Println("Too high! Try again.")
			} else {
				score = min(score, i+1)
				fmt.Printf("Correct! You guessed it in %d attempts.\n", i+1)
				fmt.Printf("Best Score: %d\n", score)

				var tryAgain string
				fmt.Print("Do you want to play again? (y/n): ")
				fmt.Scan(&tryAgain)
				if tryAgain == "y" {
					goal = rand.IntN(max) + 1
					i = -1
					continue
				} else {
					fmt.Println("Goodbye!")
					return
				}
			}
		}

		fmt.Println("You've run out of attempts! The correct number was:", goal)
		var tryAgain string
		fmt.Print("Do you want to play again? (y/n): ")
		fmt.Scan(&tryAgain)
		if tryAgain == "y" {
			goal = rand.IntN(max) + 1
			continue
		} else {
			fmt.Println("Goodbye!")
			return
		}
	}
}
