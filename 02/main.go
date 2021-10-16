package main

import "fmt"

func main() {

	var name string = "Ankit"
	var age int32 = 23
	company, address := "Google", "Delhi, India"

	fmt.Println(company)
	fmt.Println(address)
	fmt.Printf("%T %T %T", age, name, company)
}
