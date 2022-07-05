package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embed the name in a message.
	// := adalah shorcut untuk deklarasi variable, disini var message = fmt.Sprintf......
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
