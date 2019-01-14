package main

import "fmt"

/*
 * Array - Fixed-lengthh series of elements of chosen type. It is known to compile time how big they are going to be.
 * Indices start at 0 and elements of array are initialized to zero value, meaning the zero value of the type that the array is composed of.
 * So, if it's the integer, the zero value is just zero. If it's strings, the zero value is empty string. Element are accessed using subscript notation [].
 */

func main(){
	// [5] int means the array is of lenght 5 and type int.
	// so this is initialized with these values [0,0,0,0,0]
	var sampleArray [5]int

	// Assigning value to the element is by using index of the element in array
	sampleArray[0] = 2

	// Accessing element of array we use [] and index of the element in array
	fmt.Println(sampleArray[1])

	// Array Literal- An array pre-defined with values. Is the way of initializing an array to a set of values.

	var literalSample = [5]int{1,2,3,4,5}

	// ... is the key word to represent the size of the element
	// This will create array of size 5
	x := [...]int{1,2,3,4,5}

	// Iterating through the array using for loop using the range keyword
	// i(index) and v(element at index)
	for i, v := range x {
		fmt.Printf("ind %d, val %d, literalSample value %d \n", i, v, literalSample[i])
	}


}