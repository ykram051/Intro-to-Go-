package project1

import (
	"fmt"
	"math/rand/v2"
)

func v1() {
	fmt.Println("Welcome to the guessing game , try to guess a number between 0 and 100 ( inclusive ) . \n ")
	goal := rand.IntN(100)
	//fmt.Println(goal)
	var attempt int
	for {
		var userguess int
		fmt.Scan(&userguess)
		if userguess>100 || userguess<0 {
			fmt.Println("Not a valid number , Try again with a number between 0 and 100 ( inclusive)")
		}else if userguess < goal {
			fmt.Printf("Your guess is too low , Try again :")
			attempt += 1
		} else if userguess > goal {
			fmt.Printf("Your guess is too high , Try again : ")
			attempt += 1
		} else {
			fmt.Printf("Your guess is right and the number of attempts is : %d ", attempt)
			break
		}
	}

}
