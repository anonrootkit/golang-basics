package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputSize = 1000000

func main() {

	var input *bufio.Scanner
	init_input_scanner(&input)

	init_header()

	var username string
	takeInput(&username, *input)

	welcome_user(username)
}

func init_input_scanner(input **bufio.Scanner) {
	*input = bufio.NewScanner(os.Stdin)
	(*input).Buffer(make([]byte, inputSize), inputSize)
}

func init_header() {
	fmt.Print("Hey there!\n")
	fmt.Print("Let us know who you are : ")
}

func welcome_user(name string) {
	fmt.Println("Welcome, " + name)
}

func takeInput(obj *string, input bufio.Scanner) {
	input.Scan()
	fmt.Scanln(input.Text(), obj)
}
