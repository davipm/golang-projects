package main

import (
	"fmt"
	"time"
)

func main() {
	var timer1 = time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("time 1 fired")

	var timer2 = time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("time 2 fired")
	}()

	var stop2 = timer2.Stop()

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
