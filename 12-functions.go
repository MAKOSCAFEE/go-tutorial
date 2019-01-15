/*
 * @Author: Barnabas Makonda
 * @Date: 2019-01-15 22:46:35
 * @Last Modified by: Barnabas Makonda
 * @Last Modified time: 2019-01-15 22:55:54
 */

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

// variadic functions
// By using ... before the type name of the last parameter you can indicate that it takes zero or more of those parameters.
// In this case we take zero or more ints.
// We invoke the function like any other function except we can pass as many ints as we want.

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}

	return total
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	total, avarage := avarageTotal(xs)

	fmt.Println(total, avarage)

	fmt.Println(add(1, 2, 3, 4))

	// Closure - it is possible to create function within a function.
	// and it can access other local variables
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
}
