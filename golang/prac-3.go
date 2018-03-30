package main

import (
	"fmt"
	"strings"
)

func main() {
	var quote string
	fmt.Print("What is the quote? ")
	fmt.Scanln(&quote)
	quotationedQuote := []string{}
	quotationedQuote = append(quotationedQuote, "\"", quote, "\"")

	var author string
	fmt.Print("Who said it? ")
	fmt.Scanln(&author)

	fmt.Println(author, "said, ", strings.Join(quotationedQuote, ""))
}
