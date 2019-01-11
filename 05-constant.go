/*
 * @Author: Barnabas Makonda 
 * @Date: 2019-01-11 15:52:36 
 * @Last Modified by: Barnabas Makonda
 * @Last Modified time: 2019-01-11 16:23:36
 */

 package main

 import "fmt"

/*
 * Constant is an expression that is know during compile time.
 * It hold value that was declared and assigned to it all tne time as long as program is running.
 * The type is inferred from the right hand side of the expression.
 *
 *
 */


 func main() {
	 // You can assign single constant a value. Type is inferred from the right side.
	 // For this case compiler see "John" as string and hence name become of type string.
	 const name = "John"

	 // You can create multiple constants at once
	 const ( numberOne = 1 
			msg = "Jambo"
		)

	 /* iota can be used to generate set of related constant but distict. can be used to represent
	  * property that has several distinct possible values. Eg. Days of the week, Month of the years
	  * They are like enumerated(enum)type from other language.
	  */

	  type DaysOfTheWeek int
	  type Grades int

	  // you will only have to specify type once for constants and the rest will be assigned 
	  // unique value with the same datatype as the one assigned to the first one.
	  
	  const (Monday DaysOfTheWeek = iota
			Tuesday
			Wednesday
			Thursday
			Friday) 

	  const (A Grades = iota
			 B
			 C
			 D
			 E
			 F
			)

	  fmt.Println(name, numberOne, msg)
	  fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday)
	  fmt.Println(A,B,C,D,E,F)

 }