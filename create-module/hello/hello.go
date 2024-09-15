package main

import (
	"fmt"
	"log"

	"github.com/Ubunfu/godotdev/create-module/greetings"
)

func main() {
	// set properties of predefined Logger, including the log entry prefix and a flag to disable
	// printing the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message
	//message, err := greetings.Hello("Ryan")

	// Get a bunch of messages
	names := []string{"Ryan", "PJ", "Squid"}
	messages, err := greetings.Hellos(names)

	// if an error is returned, print it and exit
	if err != nil {
		// includes a baked-in call to os.Exit(1)
		log.Fatal(err)
	}

	// if no error, print the message
	fmt.Println(messages)
}
