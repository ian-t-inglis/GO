package main

import (
	"fmt"
)

func main() {
	var i int

	for i = 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizzbang")

		} else if i%5 == 0 {
			fmt.Println("Bang")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}

	}
}
