package channels

import "fmt"

// Simple example
func simple() {
	message := make(chan string)

	go func() {
		message <- "Hola soy Joshua!"
	}()

	fmt.Printf(<-message)
}
