/*
 * @Author: Barnabas Makonda
 * @Date: 2019-01-11 10:53:10
 * @Last Modified by: Barnabas Makonda
 * @Last Modified time: 2019-01-11 11:27:13
 */

package main

import "fmt"

/*
 * Variables in Go are declared explicitly.
 * Variable type is checked at the time of declaration as Go is statically typed language.
 * There are several ways of declaring and assign values to a varialbe.
 * I am going to explore them in a bit.
 */

func main() {
	// - Variable can just be declared.
	var firstNumber int // firstNumber will be set to zero
	var isPython bool   // isPython will be set to false

	// - Declare and assign

	var secondNumber = 2
	var thirdNumber = 2 // here variabble is automatically assigned as int.

	// - ShortHand definition for variable declaration.

	message := "How are you? " // message is assigned "How are you" as value and automatically assigned string as datatype.

	fmt.Println(firstNumber, isPython, secondNumber, thirdNumber, message)

	// use new() key word

	xPtr := new(int) // will return the address

	*xPtr = 1 // dereferencing it and add value to the address

	fmt.Println(*xPtr)

}
