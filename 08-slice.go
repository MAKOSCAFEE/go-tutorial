package main

import "fmt"

/*
 * Slice- A window on an underlying array. It has a variable size, up to to the whole array
 * slice can be of small or equal size of array. Three properties of slices
 *	(i) A pointer - This indicates the start of the slice
 *	(ii) Length - number of elememts in the slice
 *  (iii) Capacity - maximum number of element in the slice. This can be calculated by looking at the pointer(start of the slice) and the end of the array.
 */

func main(){
	// Whole array of size 7
	sampleArray := [...]string{"a","b","c","d","e","f", "g"}

	// First slice which start at b to d exclusive of d
	slice1 := sampleArray[1:3]

	// Second Slice which start from c to f exlclusive of f.
	slice2 := sampleArray[2:5]

	// there are two functions len() to get lenght and cap() to get capacity.

	fmt.Printf("%d , %d \n",len(slice1), cap(slice2))

	// Writing to the slice is writing to underlying array. Any change can cause changes to the other slices if there is overlaps.
	// here slice1[1] and slice2[0] referres to the same element.


	/* 
	 * Slice Literals - can be used to initialize a slice. every slice has to have an underlying array so when creating slice using slice literals,
	 * it first creates the underlying array and creates a slice to reference it. slice points to the start of the array, length is capacity.
	 * Slice will point to the entier array.
	 */


	 slice3 := []int{1,2,3}


	 fmt.Println(slice3)

	 /* A function Make(make()) can be used to create slice(and array)
	  * 2-argument version: specify type and length/capacity. length === capacity 
	  * 3-argument version: specity type, length, capacity. length !== capacity.
	  */

	  // 2-argument: type int with lenght,capacity === 10
	  slice4 := make([]int, 10)

	  // 3-argument: type int, length = 10, capacity = 15
	  // this create array of 15 element and slice of 10 element inside the array

	  slice5 := make([]int, 10,15)


	  // Append- size of a slice can be increased by adding element at the end of the array. 
	  // This increase teh size of the slice to the size of array or even increase the size of the array to fit the slice.


	  // This will add 100 at the end of slice4 and assign it to slice5
	  slice5 = append(slice4, 100)


	  fmt.Println(slice4)

	  fmt.Println(slice5)


}