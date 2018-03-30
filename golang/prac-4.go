package main

import (
	"fmt"
)

func main() {
	var noun string
	fmt.Print("Enter a noun: ")
	fmt.Scanln(&noun)

	var verb string
	fmt.Print("Enter a verb: ")
	fmt.Scanln(&verb)

	var adjective string
	fmt.Print("Enter a adjective: ")
	fmt.Scanln(&adjective)

	var adverb string
	fmt.Print("Enter a adverb: ")
	fmt.Scanln(&adverb)

	fmt.Print("Dow you ", verb, " your ", adjective, " ", noun, " ", adverb, "? That's hilarious!\n")
}
