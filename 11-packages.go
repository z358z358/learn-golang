/*
Why do we use packages?
reuse function

What is the difference between an identifier that starts with a capital letter and one which doesn’t? (Average vs average)
means other packages (and programs) are able to see it

What is a package alias? How do you make one?
rename the imported package, import f "fmt"

How would you document the functions you created in #3?
Executing: godoc golang-book/chapter11/math Min and godoc golang-book/chapter11/math Max

*/
package main

import (
	m "./packages"
	"fmt"
)

func main() {
	number := []float64{1.2, 1.1, 1.4}
	fmt.Println(m.Min(number))
	fmt.Println(m.Max(number))
}
