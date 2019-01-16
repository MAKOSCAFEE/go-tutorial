/*
 * @Author: Barnabas Makonda
 * @Date: 2019-01-15 22:46:56
 * @Last Modified by: Barnabas Makonda
 * @Last Modified time: 2019-01-16 11:06:25
 */

package main

import "fmt"

/*
 * A struct is aggregate data type. it groups together other objects of abitrary type.
 *
 */

// Person represent person
type Person struct {
	name    string
	address string
	phone   string
}

// Employee represent employee
type Employee struct {
	Person
	position string
}

// Talk is the method to print out the name or the person.
func (p *Person) Talk() {
	fmt.Println("Hi my name is ", p.name)
}

func main() {

	// Strut Person which agregate types of name, address and phone as fields of type Person

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

	// Embedded types
	mlinzi := Employee{Person{name: "Juma", address: "Mazinge st.", phone: "0277337733"}, "Mlinzi"}

	mlinzi.Talk()

}
