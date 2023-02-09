package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		secondQuery := r.URL.Query().Get("s")
		if secondQuery == "" {
			return
		}

		queryValue, err := strconv.Atoi(secondQuery)
		if err != nil {
			return
		}

		time.Sleep(time.Duration(queryValue) * time.Second)
		fmt.Printf("Responsed %v %T\n", queryValue, queryValue)
		value := fmt.Sprintf("Responsed %v", rand.Uint64())
		w.Write([]byte(value))
	})

	fmt.Println("Running Server")
	log.Fatal(http.ListenAndServe(":9999", router))
}
