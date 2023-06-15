package main

import (
	"fmt"
	"time"
)

func main() {
	// this will run forever and block the main thread from running the next line
	// count("sheep")
	// count("fish")

	// this will run concurrently
	// go count("sheep") // go keyword creates a new go routine
	// count("fish")

	// this will exit before the go routines have a chance to run
	// go count("sheep") // go immediately to the next line
	// go count("fish") // go immediately to the next line

	// this will stop when hit enter
	// go count("sheep") // go immediately to the next line
	// go count("fish") // go immediately to the next line
	// fmt.Scanln() // wait for enter
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}