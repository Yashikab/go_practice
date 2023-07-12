package main

import (
	"fmt"
	"log"

	"github.com/Yashikab/go_practice/tutorials/create_a_go_module/greetings"
)

func main() {
	// Set properties of the predefined logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("Tom")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
