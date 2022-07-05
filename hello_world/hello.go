package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote" // importing from external package
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println("Halo sop!")

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	// intinya dikasih log
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// get greeting message and print it.
	// deklare variable message dan err
	message, err := greetings.Hello("") // beri text / nama disini, ex: greetings.Hello("Sopp20")
	// Jika nilai yang dikembalikan adalah error, tampilkan di konsol
	// dan keluar dari program.
	if err != nil {
		log.Fatal(err)
	}
	// Jika tidak ada error, tampilkan message yang dikembalikan ke konsol
	fmt.Println(message)
}
