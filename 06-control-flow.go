/*
 * @Author: Barnabas Makonda 
 * @Date: 2019-01-11 12:19:39 
 * @Last Modified by: Barnabas Makonda
 * @Last Modified time: 2019-01-11 17:17:04
 */

package main

import "fmt"

/*
 * Control Flow describes the order in which statements are executed inside a program.
 * The most basic control flow is just executing one statement at a time, one after the other.
 * A procedural control flow, just top-down. A statement that alter the control flow.
 * - If statements
    if <condition> {
 *			<consequent>
 *		}
 *  Expression <condition> is evaluated. <consequent> statement(s) is only executed if <condition> is true.
 * - for loop
 *		Iterate if condition is true, may have initialization and update operation.
 * - Switch/Case
		 A multi-way if statement. contains a tag which is a variable to be checked by comparing it to a constant defined in each case.
		 case that match tag are executed.
 */

func main(){
	
	x := 5
	
	// if control statements
	if x > 6 {
		fmt.Println("More than 6")
	} else if x < 2 {
		fmt.Println("Less than 2")
	} else {
		fmt.Println("Something else")
	}

	/* for loop: 
		eg. for <init>;<cond>;<update> {
				<statement>	
			}
	*/

	for i:= 0; i < 10; i++ {
		fmt.Println(i)
	}

	// OR
	i := 100

	for i < 106 {
		fmt.Println(i)
		i++
	}


	// Switch/Case statement

	number := 2
    switch number {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
	default:
		fmt.Println(number)
	}

	// OR A tagless switch.

	switch {
	case x > 6:
		fmt.Println("More than 6")
	case x <2:
		fmt.Println("Less than 2")
	default:
		fmt.Println("Something else")
	}


	// Break and Continue. Break exits a containing for loop.
	// Continue skips the rest of the current iteration.

	// Break
	for i:= 0; i < 10; i++ {
		if i == 5 { break } 
		fmt.Println(i)  // this will only print 0 to 4 and exit the loop when i== 5
	}

	// Continue
	for i:= 0; i < 10; i++ {
		if i == 5 { continue } 
		fmt.Println(i) // this will print all the numbers except 5.
	}

	// Scan

	var ageNumber string;
	fmt.Printf("How Old are you? ")
	num, err := fmt.Scan(&ageNumber)
	fmt.Printf(ageNumber)
	fmt.Println(num, err)

	
}