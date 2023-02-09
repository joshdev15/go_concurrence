package main

import (
	"concurrence/internal/channels"
	"fmt"
)

func main() {
	fmt.Println("Inicio del Programa")
	// gosync.Run()
	// gomutex.Run()
	// race_detector.Run()
	channels.Run()
}
