package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	var x int
	x = 5
	fmt.Println(x)

	var y int = 7
	fmt.Println(y)

	var sum = x + y
	fmt.Println(sum)

	z := 10 // implicit type declaration
	fmt.Println(z)

	if z > 5 {
		fmt.Println("z is greater than 5")
	} else if z < 5 {
		fmt.Println("z is less than 5")
	} else {
		fmt.Println("z is equal to 5")
	}

	var r, t, e = 3, 4, 5
	fmt.Println(r, t, e)

	var  a [5]int
	fmt.Println(a)
	a[0] = 1
	a[1] = 2
	a[3] = 4
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	//b[7] = 6 // error: index out of range

	c := []int{1, 2, 3, 4, 5} // slice
	fmt.Println(c)
	c = append(c, 6)
	fmt.Println(c)
}