package main

import (
	"fmt"
)

func main() {
	//// for loops (only loop type in Go)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//// while loops (for loops can be used as while loops)
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	//// iterating over arrays
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	//// iterating over maps
	m := make(map[string]int)
	m["key1"] = 1
	m["key2"] = 2
	for key, value := range m { // order is not guaranteed
		fmt.Println("key:", key, "value:", value)
	}
}