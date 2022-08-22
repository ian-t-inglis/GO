package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// The GO "Rand" function has a source which is bases its numbers on.
	// In a global sense this means that every time the program is run, the same number is return
	// The following 3 lines of code define a new, local source from which Rand can generate a random number
	// This means a different random number every time the program is run.
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomNumber := r.Intn(100)

	// Local variables
	guesses := 0
	input := 0

	fmt.Println("Welcome to the random number guessing game!")
	fmt.Println("You'll have 10 guesses to find the number which has been generated from 0-100.")

	for guesses = 0; guesses < 10; guesses++ {
		fmt.Print("Enter guess: ")
		fmt.Scan(&input)

		low_close := randomNumber - input
		high_close := input - randomNumber

		if guesses == 9 {
			fmt.Println("Too many tries, better luck next time..")
		} else if input > randomNumber && high_close <= 5 {
			fmt.Println("A little high, but you're getting close")
		} else if input > randomNumber {
			fmt.Println("Too high")
		} else if input < randomNumber && low_close <= 5 {
			fmt.Println("A little low, but we're getting close")
		} else if input < randomNumber {
			fmt.Println("Too low")
		} else if input == randomNumber {
			fmt.Println("Congratulations, you guessed it!")
			fmt.Println("You've found the number in", guesses+1, "guesses!")
			break
		}

	}

}
