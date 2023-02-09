package gosync

// Para poder ver el funcionamiento debe iniciar el servidor
// que se encuentra en el directorio duration

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var (
	urlList = []string{
		"http://localhost:9999?s=3",
		"http://localhost:9999?s=1",
		"http://localhost:9999?s=5",
	}
)

func Run() {
	// secuential(urlList)
	concurrent(urlList)
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

func secuential(urls []string) {
	fmt.Println("Run Sync")

	for _, v := range urls {
		fetch(v)
	}

	fmt.Printf("\n")
}

func concurrent(urls []string) {
	fmt.Println("Run Sync")
	wg := sync.WaitGroup{}
	wg.Add(len(urls))

	for _, v := range urlList {
		go func(v string) {
			fetch(v)
			wg.Done()
		}(v)
	}

	wg.Wait()
	fmt.Printf("\n")
}
