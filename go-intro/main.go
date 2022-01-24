// EECS 491 Go Tutorial
package main

import (
	"fmt"
)

func vars() {
	// Goal: Declare an integer with value of 42
	// NOTE: In Go, fmt.Println (print line) automatically
	// inserts newline after line and a space between arguments

	/**
	 * Method 1: Explicitly state the type
	 * var <identifier> <type> = <value>
	 */
	var foo int = 42
	fmt.Println("foo's value:", foo)

	/** 
	 * Method 2: Implicit Inference
	 * Think of ':' as a hint to go for type inference
	 * <identifier> := <value>
	*/
	bar := 42
	fmt.Println("bar's value:", bar)

	/** 
	 * Pointers and Addresses (just like C++)
	*/
	ptr := &foo
	fmt.Println("ptr stores the address of foo:", ptr)

	/** 
	 * Dereferencing and assigning values
	*/
	*ptr = 69
	fmt.Println("foo now has the value:", foo)
	fmt.Println("*ptr has the same value as foo:", *ptr)
}

func cond_loops() {
	/** 
	 * Normal for-loop (with less parentheses)
	 * Go forbids pointer arithmetic
	 * i++ != ++i (++i doesn't work)
	*/
	fmt.Println("Begin conditions and loop function...")

	const END = 10
	for i := 0; i < END; i++ {
		fmt.Println("i:", i)
	}

	/** 
	 * "While loop semantics"
	 * Go doesn't have while loops, so we use for loops instead
	 * If you just have one clause in the loop, that becomes 
	 * the termination condition of the loop.
	*/
	var j int = 0
	for j < END {
		j++
	}

	/** 
	 * Conditions
	 * Similar to C++, but you can initialize values
	 * within the if-block
	*/
	if k := 9; k < 0 {
		fmt.Println("k is less than 0")
	} else if k < 10 {
		fmt.Println("k is less than 10 but greater than 0")
	} else {
		fmt.Println("k is greater than or equal to 10")
	}

	fmt.Println("End conditions and loop function...")
}

func funcs() {
	// Functions are defined with the func keyword (obv)
	fmt.Println("This is an ordinary function with no return values")
}

func addAndMultiply(a int, b int) (int, int) {
	// Functions can return multiple values
	return a + b, a * b
}

func addAndMultiplyShort(a, b int) (int, int) {
	// If an argument list has the value of the same type,
	// you don't need to specify them (although it's good to)
	return a + b, a * b
}

func main() {
	// vars()
	// cond_loops()
	// funcs()
	// a, b := addAndMultiply(1, 2)
	// c, d := addAndMultiplyShort(1, 2)
}
