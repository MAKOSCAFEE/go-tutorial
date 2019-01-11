/*
 * @Author: Barnabas Makonda 
 * @Date: 2019-01-11 11:46:17 
 * @Last Modified by: Barnabas Makonda
 * @Last Modified time: 2019-01-11 12:17:04
 */

 package main

 import "fmt"

/*
 * Scope of a variable is the place in the code where a variable can be accessed.
 * It defines how a variable referrence is resolved in the code. Scope is done using blocks(sequence of declaration and statements within matching curly bracket)
 * There are explicit block(defined by programmer using curly bracket) and implicit blocks that are not defined by blocks.
 * implicit blocks are:-
 *  - Universe block. Thats all Go source code The biggest block of all.
 *  - Package block.
 *  - File block. File Block , all the source code in a single file is within the file block,
 * 
*/

 // This be access within the file to any function.
 var x = 4

 func firstFunction() {
	 // Y is defined in a local scope and can only be accessed in firstFunction block(within firstFunction curly blacket)
	 var y = 6
	 fmt.Println(y)
	 fmt.Println(x)
 }

 func secondFunction() {

	fmt.Println(y) // Error: y is not defined.Here you can not access y because it is not in the secondFunction block and cetainly not in file block.
	fmt.Println(x)
 }

 func main() {

	firstFunction();
	secondFunction();
 }