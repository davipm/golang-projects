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

// Channel Directions
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) { fmt.Println(msg) }("going")

	time.Sleep(time.Second)
	fmt.Println("done")

	// Channels
	var message = make(chan string)
	go func() { message <- "ping" }()
	msg := <-message
	fmt.Println(msg)

	// Channel Synchronization
	var done = make(chan bool, 1)
	go worker(done)
	<-done

	// Channel Directions
	var pings = make(chan string, 1)
	var pongs = make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
