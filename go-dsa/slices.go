// Go Data Structures
package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

/**
 * Go does not have classes
 * However, we can still define methods on types
 *
 * A method is a function with a special receiver method
 * The receiver appears in its own argument between the func
 * keyword and the method name.
 *
 * In this example, the Abs method has a receiver of type Vertex named v
 */

type Vertex struct {
	X, Y float64
}

type KeyValue struct {
	Key   string
	Value string
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type sliceHeader struct {
	Length        int
	ZerothElement *byte
}

func slices() {
	/**
	 * Slice Header (passes into functions)
	 * A slice is a data structure describing a contiguous
	 * section of an array stored separately from the
	 * slice variable
	 */

	// Suppose we have an array of good ol bytes
	var buffer [256]byte

	// Idiomatic syntax
	// elements 100 to 149 inclusive of the original array
	var slice = buffer[100:150]

	// After using a slice operation on an array, we can
	// also slice a slice (elements 105 to 109 of the original array)
	slice2 := slice[5:10]

	// Reslicing
	// "slice a slice" and store the result back in the original slice
	slice2 = slice2[1:2]

	// truncate a slice
	slice = slice[1 : len(slice)-1]

	// Defining our own abstraction of the slice header
	newSlice := sliceHeader{
		Length:        5,
		ZerothElement: &buffer[105],
	}

	fmt.Println(slice)
	fmt.Println(newSlice)
	fmt.Println(slice2)
}

func WCMap(value string) []KeyValue {
	mapping := make([]KeyValue, 0)

	return mapping
}

func main() {
	var text string = "THE KING JAMES BIBLE - PUBLIC DOMAIN\n\n         \n         From: Bill McGinnis Ministries - \"Feeding His Sheep\"\n         http://www.patriot.net/users/bmcgin/ministries.html\n             P. O. Box 2543, Alexandria, VA 22301   USA\n                 bmcgin@patriot.net - (703)-768-6710\n  \n                     \n                         February 19, 2000\n                            May 22, 2002\n\nHere is an excellent Public Domain version of the King James\nBible (KJV).\n\nThe official statement which places it in the Public Domain\nis shown here . . .\n\n****************************************************************\nThis version of the King James Bible was created by taking several\npublic domain copies and painstakingly comparing them to find all\nthe differences, and selecting the most common version. Each of the\ndifferences was also compared to printed versions for verification."
	functor := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	words := strings.FieldsFunc(text, functor)
	for i, word := range words {
		fmt.Println("index:", i, "word:", word)
	}
}
