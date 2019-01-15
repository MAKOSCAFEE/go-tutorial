/*
 * @Author: Barnabas Makonda
 * @Date: 2019-01-15 23:34:13
 * @Last Modified by: Barnabas Makonda
 * @Last Modified time: 2019-01-15 23:41:24
 */

package main

import "fmt"

func isEven(number int) bool {
	return number%2 == 0
}

func halfAndCheck(number int) (int, bool) {
	halfANumber := number / 2

	isEvenNumber := isEven(number)

	return halfANumber, isEvenNumber
}

func main() {
	half, isEven := halfAndCheck(1)

	fmt.Println(half, isEven)
}
