package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name string
	age  int
}

func (human Human) greet() {
	fmt.Println("Hello my name is " + human.name + " and I'm " + strconv.Itoa(human.age) + " years old.")
}

func (human *Human) hasBirthday() {
	human.age = human.age + 1
}

func main() {
	ankit := Human{name: "Ankit", age: 23}

	ankit.greet()
	ankit.hasBirthday()
	ankit.greet()
}
