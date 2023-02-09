package channels

import (
	"fmt"
	"time"
)

type account struct {
	name    string
	balance int
}

type backOperation struct {
	amount int
	done   chan struct{}
}

func CSPAndTrad() {
	fmt.Println("Run CSP and Traditional Concurrence")
	transactions := []int{300, 300}

	signal := make(chan struct{})
	transaction := make(chan *backOperation)

	katty := account{name: "Katty", balance: 500}
	joshua := account{name: "Joshua", balance: 900}

	go func() {
		for {
			request := <-transaction
			transfer(request.amount, &katty, &joshua)
			request.done <- struct{}{}
		}
	}()

	for _, value := range transactions {
		go func(amount int) {
			requestTransaction := backOperation{amount: amount, done: make(chan struct{})}
			transaction <- &requestTransaction
			<-requestTransaction.done
			signal <- struct{}{}
		}(value)
	}

	<-signal
	<-signal
}

func transfer(amount int, source, dest *account) {
	if source.balance < amount {
		fmt.Printf("❌: %v\n", fmt.Sprintf("%v %v -> %v %v", source.name, source.balance, dest.name, dest.balance))
		return
	}

	time.Sleep(time.Second)

	dest.balance += amount
	source.balance -= amount

	fmt.Printf("✅: %v\n", fmt.Sprintf("%v %v -> %v %v", source.name, source.balance, dest.name, dest.balance))
}
