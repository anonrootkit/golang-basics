package main

import "fmt"

func main() {
	a := 50
	c := &a

	fmt.Printf("%T %d \n", a, **&c)

	*c = 100

	fmt.Printf("%T %d \n", a, **&c)
}
