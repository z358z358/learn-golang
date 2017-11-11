// Write a series of tests for the Min and Max functions you wrote in the previous chapter.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(Average([]float64{}))

}

// Writing a good suite of tests is not always easy, but the process of writings tests often reveals more about a problem then you may at first realize. For example, with our Average function what happens if you pass in an empty list ([]float64{})? How could we modify the function to return 0 in this case?
func Average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
