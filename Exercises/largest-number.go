/*
 * @Author: Barnabas Makonda
 * @Date: 2019-01-15 23:21:23
 * @Last Modified by: Barnabas Makonda
 * @Last Modified time: 2019-01-15 23:32:46
 */
package main

import "fmt"

/* Question:
 * Write a function with one variadic parameter that finds the greatest number in a list of numbers.
 *
 */

func largestNumber(args ...int) int {
	largest := args[0]

	for _, value := range args {
		if value > largest {
			largest = value
		}
	}

	return largest
}

func main() {
	fmt.Println(largestNumber(2, 68, 38, 90, 8))
}
