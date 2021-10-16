package main

import "fmt"

func main() {
	// for loops

	i := 1
	for i < 10 { // for i := 0; i < 10; i++
		fmt.Printf("%d\t", i)
		i++
	}

	fmt.Print("\n\n\n")

	// Fizzbuzz
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz \t")
		} else if i%3 == 0 {
			fmt.Print("Fizz \t")
		} else if i%5 == 0 {
			fmt.Print("Buzz \t")
		}
	}
}
