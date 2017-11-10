package main

import "fmt"

// sum is a function which takes a slice of numbers and adds them together. What would its function signature look like in Go?
// func sum(args ...int)
func main() {
	h1, b1 := half(1)
	h2, b2 := half(2)
	fmt.Println("half(1)", h1, b1, "half(2)", h2, b2)

	fmt.Println("greatest number in a list of numbers:", greatest(48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17))

	nextEven := makeOddGenerator()
	fmt.Println(nextEven()) // 1
	fmt.Println(nextEven()) // 3
	fmt.Println(nextEven()) // 5

	fmt.Println("fib(2)", fib(2), "fib(3)", fib(3), "fib(10)", fib(10))

	// What are defer, panic and recover? How do you recover from a run-time panic?
	defer func() {
		str := recover()
		fmt.Println(str, "hi")
	}()
	panic("PANIC")
}

// Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).
func half(x int) (r int, even bool) {
	r = x / 2
	even = x%2 == 0
	return
}

// Write a function with one variadic parameter that finds the greatest number in a list of numbers.
func greatest(args ...int) (r int) {
	for _, value := range args {
		if r == 0 || value > r {
			r = value
		}
	}
	return
}

// Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers.
func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
