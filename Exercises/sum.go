/*
 * @Author: Barnabas Makonda
 * @Date: 2019-01-15 23:15:36
 * @Last Modified by: Barnabas Makonda
 * @Last Modified time: 2019-01-15 23:20:22
 */

package main

import "fmt"

/* Question:
 * sum is a function which takes a slice of numbers and adds them together.
 * What would its function signature look like in Go?
 */

func sum(args []int) int {
	total := 0
	for _, value := range args {
		total += value
	}
	return total

}

func main() {
	slice1 := []int{2, 3, 4, 5, 6}

	fmt.Println(sum(slice1))
}
