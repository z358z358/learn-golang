// How do you get the memory address of a variable?
// &variable

// How do you assign a value to a pointer?
// *pointer = 1

// How do you create a new pointer?
// pointer := new(int)

package main

import "fmt"

func main() {
	xx := 1.5
	square(&xx)

	fmt.Println(xx)

	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}

// What is the value of x after running this program:
// 2.25
func square(x *float64) {
	*x = *x * *x
}

// Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).
func swap(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}
