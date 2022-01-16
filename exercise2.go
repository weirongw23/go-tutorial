// Exercise 2
// Structs, Maps, Slices, panic (exceptions), make keyword (like new)
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func scan(filename string, counts map[string]int) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		words := strings.Fields(line)

		for _, w := range words {
			counts[w]++
		}
	}

	return nil
}

func main() {
	counts := make(map[string]int)

	for _, p := range os.Args[1:] {
		if err := scan(p, counts); err != nil {
			log.Fatal(err)
		}
	}

	for w, c := range counts {
		fmt.Printf("word: %s | count: %d\n", w, c)
	}
}