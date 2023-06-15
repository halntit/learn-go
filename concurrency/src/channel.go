package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string) // create a channel of type string
	go count("sheep", c)

	// for {
	// 	msg, open := <-c
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	for msg := range c {
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) { // c: channel of strings (c is a channel type and channel also has a type, which is string)
	for i := 1; i <= 5; i++ {
		c <- thing // send a string to the channel
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
