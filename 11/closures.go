package main

import "fmt"

func adder() func(int) int {
	sum := 0

	return func(num int) int {
		sum += num
		return sum
	}
}

func main() {
	name := "Ankit"

	var namePrinter = func(name string) string {
		return name
	}

	fmt.Println(namePrinter(name))

	var adder_func = adder()

	fmt.Println(adder_func(10))
	fmt.Println(adder_func(10))
}
