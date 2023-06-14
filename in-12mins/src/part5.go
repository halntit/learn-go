package main

import (
	"fmt"
)

func main() {
	i := 7
	fmt.Println(i)
	fmt.Println(&i) // & gives you the memory address of the variable (pointer)

	inc(i) // pass by value, i is copied into the function
	fmt.Println(i)

	inc2(&i) // pass by reference, i is not copied into the function
	fmt.Println(i)
}

func inc(x int) {
	x++
}

func inc2(x *int) { // pass by reference, x is a pointer to the variable
	*x++
}