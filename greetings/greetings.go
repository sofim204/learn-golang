package greetings

import (
	"errors"
	"fmt"
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

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
