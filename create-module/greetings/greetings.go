package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name, retturn an error message
	if name == "" {
		return "", errors.New("empty name")
	}

	// if name provided, return a greeting embedded with the name
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// a map to associate with messages
	messages := make(map[string]string)

	// loop through the names, getting a greeting for each one
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

// ramdomFormat returns a greeting message format chosen at random
func randomFormat() string {
	// the slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// return a random format
	return formats[rand.Intn(len(formats))]
}
