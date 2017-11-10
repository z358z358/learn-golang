/*
Problems

How do you access the 4th element of an array or slice?
array[3]

What is the length of a slice created using: make([]int, 3, 9)?
3

Given the array:

x := [6]string{"a","b","c","d","e","f"}
what would x[2:5] give you?
{"c","d","e"}

*/

package main

import "fmt"

func main() {
	// What is the length of a slice created using: make([]int, 3, 9)?
	//3
	fmt.Println(len(make([]int, 3, 9)))

	//Given the array:

	x := [6]string{"a", "b", "c", "d", "e", "f"}
	//what would x[2:5] give you?
	//{"c","d","e"}

	fmt.Println(x[2:5])
	// Write a program that finds the smallest number in this list:

	xx := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	small := 0

	for _, value := range xx {
		if small == 0 {
			small = value
		} else if value < small {
			small = value
		}
	}
	fmt.Println("smallest number in this list:", small)

}
