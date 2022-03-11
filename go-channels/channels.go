package main

import "fmt"

func main() {
	// To create a new channel (buffer = 2) called messages
	messages := make(chan string, 2)

	// Create a goroutine and send a message to the channel
	go func() { messages <- "hello from goroutine 1" }()

	// Create a second goroutine
	go func() { messages <- "hello from goroutine 2" }()

	// Get the message from the channel
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
