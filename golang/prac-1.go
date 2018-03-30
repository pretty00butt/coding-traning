package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("What is your name? ")
	fmt.Scan(&input)
	fmt.Println("Hello,", input)
}
