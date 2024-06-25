package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// Calling Goroutine
	go display("Welcome", start)
}

func display(message string, start time.Time) {
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(message)
	}
	timeElapsed := time.Since(start).Seconds()
	fmt.Printf("The `for` loop took %v", timeElapsed)
}
