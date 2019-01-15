package main

import "fmt"

/*
 * A struct is aggregate data type. it groups together other objects of abitrary type.
 *
 */

func main() {

	// Strut Person which agregate types of name, address and phone as fields of type Person

	type Person struct {
		name    string
		address string
		phone   string
	}

	var tzPresident Person

	// Accessing the struct fields, we use dot notation
	tzPresident.name = "John Pombe Magufuli"
	tzPresident.address = "Magogoni street"
	tzPresident.phone = "(222) 0 1234445"

	x := tzPresident.name

	fmt.Println(x)

	// Initializing struct we can use new() which will initialize all fields to zero data of that field type. if int, then 0 if string then ""
	p2 := new(Person)

	fmt.Println(p2)

	// OR initialize using struct literal

	p1 := Person{name: "joe", address: "Mazinge st.", phone: "0277337733"}

	fmt.Println(p1)

}
