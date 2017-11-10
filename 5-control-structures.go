package main

import "fmt"

func main() {
	//What does the following program print:
	// Small

	i := 10
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}

	//Write a program that prints out all the numbers evenly divisible by 3 between 1 and 100. (3, 6, 9, etc.)
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	//Write a program that prints the numbers from 1 to 100. But for multiples of three print "Fizz" instead of the number and for the multiples of five print "Buzz". For numbers which are multiples of both three and five print "FizzBuzz".
	for i := 1; i <= 100; i++ {
		tag := false
		if i%3 == 0 {
			fmt.Print("Fizz")
			tag = true
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
			tag = true
		}

		if tag {
			fmt.Println("")
		}
	}
}
