package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote" // importing from external package (https://pkg.go.dev/search?q=quote)
)

func main() {
	fmt.Println(quote.Glass()) // pilihan function nya ada quote.Hello(), quote.Glass(), quote.Go(), quote.Opt()
	fmt.Println("Halo sop!")

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	// intinya dikasih log
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// get greeting message and print it.
	// deklare variable message dan err
	message, err := greetings.Hello("Muhammad Sofi") // beri text / nama disini, ex: greetings.Hello("Sopp20")
	// Jika nilai yang dikembalikan adalah error, tampilkan di konsol
	// dan keluar dari program.
	if err != nil {
		log.Fatal(err)
	}
	// Jika tidak ada error, tampilkan message yang dikembalikan ke konsol
	fmt.Println(message)
}
