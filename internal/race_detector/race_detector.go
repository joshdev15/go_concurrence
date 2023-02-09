package race_detector

import (
	"fmt"
	"sync"
)

// Comment mutex variable in all file and run "go run -race main.go"
// to show the data race warning message
func Run() {
	mutex := sync.Mutex{}
	fmt.Println("Race Detector")

	data := 1

	go func() {
		mutex.Lock()
		data++
		mutex.Unlock()
	}()

	mutex.Lock()
	fmt.Printf("Data: %v\n", data)
	mutex.Unlock()
}
