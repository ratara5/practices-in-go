/*
GO TUTORIAL: 1.GETTING STARTED and 2.(partial) CREATE A GO MODULE 3.(partial) RETURN AND HANDLE ERROR
*/
package main

import "fmt"
import "rsc.io/quote"
import "example.com/greetings"
import "log"

func main() {
    fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names
    names := []string{"Raynet", "Ricky", "Alvar"}

    // Request a greeting message.
    // message, err := greetings.Hello("")
    // message, err := greetings.Hello("Raynet")

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)

    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
	// 2.(partial) Get a greeting message and print it.
    // fmt.Println(message)

    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)
}