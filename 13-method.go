/*
 * @Author: Barnabas Makonda
 * @Date: 2019-01-16 10:11:27
 * @Last Modified by: Barnabas Makonda
 * @Last Modified time: 2019-01-16 10:19:48
 */

package main

/*
 * Methods are like functions.
 * In between the keyword func and the name of the function we've added a “receiver”.
 * The receiver is like a parameter – it has a name and a type – but by creating the function in this way it allows us to call the function using the
 *
 */

import "fmt"

// Rectangle : takes for corners to get width and length.
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}
