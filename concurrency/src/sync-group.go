package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup // create a wait group variable (counter)
	wg.Add(1) // add 1 to the wait group counter
	go count("sheep", &wg)

	wg.Wait() // block the main thread until the wait group is done
}

func count(thing string, wg *sync.WaitGroup) { // wg is a pointer to a wait group
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}

	wg.Done() // decrement the wait group counter
}