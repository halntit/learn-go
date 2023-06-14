package main

import (
	"fmt"
	"errors"
	"math"
)

// type uses to create a new type, in this case a struct
type Person struct {
	name string
	age int
}

func main() {
	person1 := Person{name: "John", age: 25}
	fmt.Println(person1)
	fmt.Println(person1.name)
	fmt.Println(person1.age)

	result := sum(2, 5)
	fmt.Println(result)

	result2 := sqrt(16)
	fmt.Println(result2)

	result3, err := sqroot(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result3)
	}
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func sqroot(x float64) (float64, error) { // return multiple values
	if x < 0 { // error handling, Go does not have exceptions or try-catch blocks
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}