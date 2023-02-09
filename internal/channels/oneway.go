package channels

import "fmt"

// Oneway example
func oneway() {
	number := make(chan int)
	go receive(number)
	send(number)
}

// Define the operation that can be performed on the channel,
// at the time of signing the function

func send(number chan<- int) {
	number <- 10
}

func receive(number <-chan int) {
	fmt.Println(<-number)
}
