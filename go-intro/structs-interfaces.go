package structsInterfaces

import (
	"fmt"
	"math"
)

/**
 * PART I: Structs and has-a Relationship
 */

/**
 * Circle Class & Methods
 */
type Circle struct {
	x float64
	y float64
	r float64
}

// change the circle to pass-by-value
// This is a method because (c *Circle) is a receiver
func (c *Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func circleFunc() {
	// Initialize an instance of circle locally
	c1 := Circle{x: 0, y: 0, r: 5}

	// Initialize an instance of circle by allocating memory
	// and set each field to corresponding 0 value
	// This returns a pointer (since it's allocated on the heap)
	c2 := new(Circle)

	fmt.Println(c1)
	fmt.Println(c2)

	c1.x = 10
	c1.y = 5
	c1.r = 7

	fmt.Println("The area of circle 1 is: ", c1.Area())
}

/** 
 * Rectangle Class and Methods
*/
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) Area() float64 {
	length := distance(r.x1, r.y1, r.x1, r.y2)
	width := distance(r.x1, r.y1, r.x2, r.y1)
	return length * width
}

/** 
 * PART II: Embedded Types, Interfaces, is-a Relationship
*/

// Embedded Type (anonymous fields)
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

// Since an Android is a Person,
// and a Person can talk,
// Android can also talk
func AndroidPerson() {
	a := new(Android)
	a.Person.Talk()

	// OR a.Talk()
}

// Defining Interfaces: type name interface
// Here, we define a method set, a list of
// methods that a type must have in order
// to "implement" the interface.
type Shape interface {
	Area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, shape := range shapes {
		area += shape.Area()
	}

	return area
}

// Interfaces can also be used as fields
type MultiShape struct {
	shapes []Shape
}

// Turn Multishape into a shape by giving it an area method
func (m *MultiShape) Area() float64 {
	var area float64
	for _, shape := range m.shapes {
		area += shape.Area()
	}
	return area
}