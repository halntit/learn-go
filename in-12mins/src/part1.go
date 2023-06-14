package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	//// declaring variables
	var x int // explicit type declaration
	x = 5
	fmt.Println(x)

	var y int = 7
	fmt.Println(y)

	var sum = x + y
	fmt.Println(sum)

	z := 10 // implicit type declaration
	fmt.Println(z)

	//// if else statements
	if z > 5 {
		fmt.Println("z is greater than 5")
	} else if z < 5 {
		fmt.Println("z is less than 5")
	} else {
		fmt.Println("z is equal to 5")
	}

	//// declaring multiple variables at once
	var r, t, e = 3, 4, 5
	fmt.Println(r, t, e)

	//// declearing arrays (fixed size)
	var a [5]int
	fmt.Println(a)
	a[0] = 1
	a[1] = 2
	a[3] = 4
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	//b[7] = 6 // error: index out of range

	//// declaring slices (dynamic size)
	c := []int{1, 2, 3, 4, 5}
	fmt.Println(c)
	c = append(c, 6) // append returns a new slice
	fmt.Println(c)
}