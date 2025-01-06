package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	fmt.Println("Welcome to the guessing game . \n ")
	var Numtries int
	fmt.Println("Enter the maximum number of tries : ")
	fmt.Scan(&Numtries)
	var score int = Numtries

	var difficulty int
	fmt.Println("Enter Wanted difficulty 0 for easy: 1-50 , 1 for Medium: 1-100 ,2 for Hard: 1-200")
	for {
		fmt.Scan(&difficulty)
		if difficulty != 0 && difficulty != 1 && difficulty != 2 {
			fmt.Println("The difficulty chosen must be either 0,1 or 2 Try again! ")
		} else {
			fmt.Printf("Difficulty selected is : %d", difficulty)
			break
		}
	}
	var max int
	var goal int
	switch difficulty {
	case 0:
		goal = rand.IntN(50)
		max = 50
	case 1:
		goal = rand.IntN(100)
		max = 100
	case 2:
		goal = rand.IntN(200)
		max = 200
	}

	//fmt.Println(goal)

	for i := 0; i < Numtries+1; i++ {
		var userguess int
		fmt.Println("Enter a number with the chosen range ")
		fmt.Scan(&userguess)
		if userguess > max && i < Numtries || userguess < 0 && i < Numtries {
			fmt.Printf("Not a valid number , Try again with a number between 0 and %d ( inclusive)", max)
			i--
		} else if userguess < goal && i < Numtries {
			fmt.Printf("Your guess is too low , Try again : ")
		} else if userguess > goal && i < Numtries {
			fmt.Printf("Your guess is too high , Try again : ")
		} else if userguess == goal {
			score = min(score, i)
			fmt.Printf("Your guess is right and the number of attempts now is : %d \n ", i)
			fmt.Printf("Your highest score now is : %d ", score)
			var tryAgain string
			fmt.Print("Do you want to try again ? enter y for yes and n for no ")
			fmt.Scan(&tryAgain)
			switch tryAgain {
			case "y":
				i = 0
			case "n":
				fmt.Printf("GoodBye")

			}

		} else if i == Numtries {
			var tryAgain string
			fmt.Printf("Your number of guesses is up , You lost , Do you want to try again ? enter y for yes and n for no ")
			fmt.Scan(&tryAgain)
			switch tryAgain {
			case "y":
				i = 0
			case "n":
				fmt.Printf("GoodBye")

			}
		}
	}

}
