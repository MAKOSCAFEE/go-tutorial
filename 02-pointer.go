/*
 * @Author: Barnabas Makonda 
 * @Date: 2019-01-11 10:13:57 
 * @Last Modified by: Barnabas Makonda
 * @Last Modified time: 2019-01-11 10:50:12
 */


package main

import "fmt"

/*
 A pointer is an address to some data in memory.
 Variable, functions, etc are all located in memory somewhere.
 A pointer is the address of that in memory.

 There are two main operation associated with the pointer
	- The ampersand(&) operator that return the address of the variable,functions,etc.
	  If put infront of the variable, name of the variable it will return address of that variable.
	- The star(*) operator which dereferrencing and returns the data at that address. If put infront
	  of the pointer to some address it will return the data at that address.
*/

func main() {
	var x int  = 1
	var y int

	var ip *int // variable ip is the pointer to an interger

	ip = &x // ip now points to x, so ip holds the address to x.
	y = *ip // y is now 1 as it derefferrence the pointer ip and return the data at address ip.

	fmt.Println("Value of y is ", y)
	
}

