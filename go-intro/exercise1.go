package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func scan(filename string) error {
	// Common paradigm in Go
	// os.Open(filename) returns (file *, error)
	// file * is the file descriptor
	// if error is encountered, it returns *PathError
	f, err := os.Open(filename)

	// In Go, nil contains the value 0
	// 0 means everything is good
	// Any other value than 0 means a problem is encountered
	if err != nil {
		return err
	}

	// defer keyword ensures this happens no matter the execution
	// exceptions, abort, etc. can ensure resources are freed
	// if there are no error, this closes at the end so you don't forget it
	// useful for preventing deadlocking (lock(), defer unlock()), like a dtor
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Use a "while" loop
	for scanner.Scan() {
		// What is the index of the word "the"
		// Read each line individually, split the words
		// into an array of sorts, then determine the index of "the"
		line := scanner.Text()
		words := strings.Fields(line)

		// range allows us to iterate through any iterable object
		// range gives an index and the actual object
		for i, word := range words {
			if word == "the" {
				fmt.Printf("%s is the %d th word\n", word, i)
			}
		}
	}

	return scanner.Err()
}

func main() {
	for _, filename := range os.Args[1:] {
		if err := scan(filename); err != nil {
			log.Fatal(err)
		}
	}
}