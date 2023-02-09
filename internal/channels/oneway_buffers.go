package channels

import (
	"fmt"
	"time"
)

// Oneway with Buffer example
func onewayBuffers() {
	number := make(chan int, 2)
	signal := make(chan struct{})

	go receiveBuffer(signal, number)
	sendBuffer(number)

	signal <- struct{}{}
}

// Define the operation that can be performed on the channel,
// at the time of signing the function

func sendBuffer(number chan<- int) {
	number <- 1
	number <- 2
	number <- 3
	number <- 4
	number <- 5
	time.Sleep(time.Nanosecond)
	number <- 6

	// with close
	// close(number)
}

func receiveBuffer(signal <-chan struct{}, number <-chan int) {
	// Without Close
	// fmt.Println(<-number)
	// fmt.Println(<-number)
	// fmt.Println(<-number)

	// With Close

	// for range
	// for v := range number {
	// fmt.Printf("%d \n", v)
	// }

	// for select
	for {
		select {
		case v := <-number:
			fmt.Println(v)

		case <-signal:
			return

		default:
			fmt.Println("🤔")
		}
	}
}
