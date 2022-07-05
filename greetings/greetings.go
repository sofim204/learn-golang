package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// menambahkan error handling //
// Hello return a greeting for the named person.
func Hello(name string) (string, error) {
	// Return a greeting that embed the name in a message.
	// := adalah shorcut untuk deklarasi variable, disini var message = fmt.Sprintf......

	// Jika tidak ada nama yang diberikan, beri return error berupa pesan error.
	if name == "" {
		return "", errors.New("tidak ada nama yang diberikan")
	}

	// Jika nama diterima, kembalikan value yang sudah diberikan
	// di greeting message.

	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name) // part of random greetings
	return message, nil
}

// random greetings start here //
// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set greeting message.
// the returned message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Nice to meet you, %v. Let's go to Park next monday.",
	}

	// Return a random selected message format by specifying
	// a random index fro the slice of formats.
	return formats[rand.Intn(len(formats))]
}
