package main

import "fmt"

func main() {
	//Defining the variables to be used for calculation.
	//The equation for the Fibonacci sequence is Xn = Xn-1 + Xn-2
	//Or in simpler terms, the current number is equal to the previous number added to the number before the previous number
	//In this case, i = Xn, j = Xn-1, and k=Xn-2
	//iter and input are used purely for the loop
	i := 1
	j := 1
	k := 1
	iter := 0
	input := 0

	//This section here takes input from the user
	fmt.Print("Enter a positive number: ")
	fmt.Scan(&input)

	//This for loop will constantly loop back if the input entered is not a positive integer
	for input <= 0 {
		fmt.Println("Please enter a positive number, negatives will not work.")
		fmt.Print("Enter number: ")
		fmt.Scan(&input)
	}

	//This is where the fibonacci equation is run to return a number along the sequence
	for iter = 0; iter < input; iter++ {
		k = j
		j = i
		i = k + j

	}
	fmt.Println("The fibonacci number after", input, "iterations is", i)
}
