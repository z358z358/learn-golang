package main

import (
	"fmt"
)

/*
Problems

What is whitespace?
Go mostly doesn't care about whitespace

What is a comment? What are the two ways of writing a comment?
// and /*

Our program began with package main. What would the files in the fmt package begin with?
fmt

We used the Println function defined in the fmt package. If we wanted to use the Exit function from the os package what would we need to do?
import "os"

*/

func main() {
	//Modify the program we wrote so that instead of printing Hello World it prints Hello, my name is followed by your name.
	fmt.Println("Hello my name is Onion")
}
