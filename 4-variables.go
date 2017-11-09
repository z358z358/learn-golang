/*

Problems

What are two ways to create a new variable?
var x int
x := 0

What is scope and how do you determine the scope of a variable in Go?
inside or outside of the function

What is the difference between var and const?
const cant change

Write another program that converts from feet into meters. (1 ft = 0.3048 m)
*/
package main

import "fmt"

func main() {
	// What is the value of x after running: (x := 5; x += 1?
	// 6
	x := 5
	x += 1
	fmt.Println(x)

	// Using the example program as a starting point, write a program that converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)
	f := 90
	c := (f - 32) * 5 / 9
	fmt.Println(f, c)

	// Write another program that converts from feet into meters. (1 ft = 0.3048 m)
	feet := 1.0
	meters := feet * 0.3048
	fmt.Println(feet, meters)
}
