package gomutex

import (
	"fmt"
	"sync"
	"time"
)

type account struct {
	name    string
	balance int
}

func Run() {
	// mutexSecuential()
	mutexConcurrent()
}

func mutexConcurrent() {
	fmt.Println("Run Mutex Concurrent")

	transactions := []int{300, 300}

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	wg.Add(len(transactions))

	katty := account{name: "Katty", balance: 500}
	joshua := account{name: "Joshua", balance: 900}

	for _, value := range transactions {
		go func(amount int) {
			mutex.Lock()
			transfer(amount, &katty, &joshua)
			mutex.Unlock()
			wg.Done()
		}(value)
	}

	wg.Wait()
}

func mutexSecuential() {
	fmt.Println("Run Mutex Secuential")

	transactions := []int{300, 300}

	katty := account{name: "Katty", balance: 500}
	joshua := account{name: "Joshua", balance: 900}

	for _, value := range transactions {
		transfer(value, &katty, &joshua)
	}

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
