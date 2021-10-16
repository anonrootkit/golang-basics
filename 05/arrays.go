package main

import (
	"fmt"
)

func main() {

	//Arrays
	var names [3]string
	var ages [3]int

	names[0], names[1], names[2] = "Ankit", "Vipin", "Tripathi"
	ages[0], ages[1], ages[2] = 23, 20, 24

	appleSlices := []int{1, 2, 3, 4, 5}

	fmt.Println(len(appleSlices))
	fmt.Println(appleSlices[0:4])

}
