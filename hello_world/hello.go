package main

import (
	"fmt"

	"example.com/greetings"
	"rsc.io/quote" // importing from external package
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println("Halo sop!")
	// get greeting message and print it.
	message := greetings.Hello("Muhammad Sofi")
	fmt.Println(message)
}
