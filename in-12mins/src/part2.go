package main

import (
	"fmt"
)

func main() {
	//// declaring maps
	var m map[string]int
	fmt.Println(m)

	m = make(map[string]int) // initialize map, otherwise it will be nil
	fmt.Println(m)

	m["key1"] = 1
	m["key2"] = 2
	fmt.Println(m)

	//// shorthand for declaring and initializing maps
	m2 := make(map[string]int)
	m2["key1"] = 1
	m2["key2"] = 2
	fmt.Println(m2)

	//// declaring and initializing maps in one line
	vertices := map[string]int{
		"triangle": 3,
		"square": 4,
		"dodecagon": 12,
	}
	fmt.Println(vertices)
	fmt.Println(vertices["triangle"]) // accessing value of a key

	delete(vertices, "dodecagon") // deleting a key
	fmt.Println(vertices)

	delete(vertices, "pentagon") // deleting a non-existent key, no error
	fmt.Println(vertices)
}