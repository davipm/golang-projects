package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// Channel Synchronization
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) { fmt.Println(msg) }("going")

	time.Sleep(time.Second)
	fmt.Println("done")

	// Channels
	message := make(chan string)
	go func() { message <- "ping" }()
	msg := <-message
	fmt.Println(msg)

	// Channel Synchronization
	done := make(chan bool, 1)
	go worker(done)
	<-done
}
