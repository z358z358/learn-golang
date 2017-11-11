// We copied the average function from chapter 7 to our new package. Create Min and Max functions which find the minimum and maximum values in a slice of float64s.

package math

func Min(numbers []float64) (r float64) {

	for _, value := range numbers {
		if r == 0 || r > value {
			r = value
		}
	}
	return
}

func Max(numbers []float64) (r float64) {

	for _, value := range numbers {
		if r == 0 || r < value {
			r = value
		}
	}
	return
}
