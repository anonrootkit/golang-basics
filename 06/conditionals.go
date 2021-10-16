package main

import "fmt"

func main() {
	x, y := 10, 20

	if x > y {
		fmt.Print("X is greater")
	} else {
		fmt.Print("Y is greater")
	}

	switch x {
	case 10:
		fmt.Printf("Number is 10")
	case 20:
		fmt.Print("Number is 20")
	}
}
