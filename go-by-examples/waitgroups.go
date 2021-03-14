package main

import (
	"fmt"
	"sync"
	"time"
)

func worker2(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %v starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %v done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker2(i, &wg)
	}

	wg.Wait()
}
