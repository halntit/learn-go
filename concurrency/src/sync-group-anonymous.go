package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup // create a wait group variable (counter)
	wg.Add(1) // add 1 to the wait group counter

	go func() { // an anonymous goroutine
		count("sheep")
		wg.Done()
	}()

	wg.Wait() // block the main thread until the wait group is done
}

func count(thing string) { // not changing the function signature
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}