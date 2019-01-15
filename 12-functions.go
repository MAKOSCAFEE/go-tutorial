package main

import "fmt"

/*
 *Function - an independent section of code that maps zero or more input parameters to zero or more output parameters.
 */

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// Return multiple values

func avarageTotal(xs []float64) (float64, float64) {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total, total / float64(len(xs))
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	total, avarage := avarageTotal(xs)

	fmt.Println(total, avarage)
}
