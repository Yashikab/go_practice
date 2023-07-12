package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	//If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
