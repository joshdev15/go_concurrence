package channels

// Para poder ver el funcionamiento debe iniciar el servidor
// que se encuentra en el directorio duration

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	urlList = []string{
		"http://localhost:9999?s=3",
		"http://localhost:9999?s=2",
		"http://localhost:9999?s=6",
		"http://localhost:9999?s=1",
		"http://localhost:9999?s=4",
		"http://localhost:9999?s=7",
	}
)

func syncToCSP() {
	// signalExample(urlList)
	withCancelation(urlList)
}

func fetch(value string) {
	response, err := http.Get(value)
	if err != nil {
		text := err.Error()
		fmt.Printf("%v\n", text)
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		text := err.Error()
		fmt.Printf("%v\n", text)
	}

	fmt.Printf("%v\n", string(bytes))
}

func signalExample(urls []string) {
	fmt.Println("Run SyncToCSP")
	signal := make(chan struct{})

	for _, v := range urlList {
		go func(v string) {
			fetch(v)
			signal <- struct{}{}
		}(v)
	}

	<-signal
	<-signal
	<-signal
	fmt.Printf("\n")
}

func withCancelation(urls []string) {
	fmt.Println("Run With Cancel")
	signal := make(chan struct{})

	for _, v := range urlList {
		go func(v string) {
			fetch(v)

			select {
			case <-signal:
				return
			}
		}(v)
	}

	select {
	case <-time.After(time.Second * 5):
		close(signal)
	}

	fmt.Printf("\n")
}
