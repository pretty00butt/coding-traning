package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("What is your name? ")
	fmt.Scanln(&input)
	if len(input) == 0 {
		fmt.Println("input is empty!")
	} else {
		fmt.Println(input, " has ", len(input), "characters. ")
	}
}
