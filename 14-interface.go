/*
 * @Author: Barnabas Makonda
 * @Date: 2019-01-16 11:11:52
 * @Last Modified by: Barnabas Makonda
 * @Last Modified time: 2019-01-16 11:22:03
 */

package main

import (
	"fmt"
	"math"
)

/*
 * Like struct an interface is created using the type keyword followed by the name and the keyword interface.
 * But instead of defining fields we  define a "method set". A method set is a list of methods that a type must have in order to implement the interface.
 * They can be used as field type too.
 */

// Circle represent a circle
type Circle struct {
	x, y, r float64
}

// Rectangle struct
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Shape Interface with a method
type Shape interface {
	area() float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// All argument must implement method area as interface Shape suggest.
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}

	fmt.Println(totalArea(&c, &r))
}
