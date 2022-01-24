// Exercise 3 - Concurrency
// Goroutines and Channels (thread safe)
// For collections, Go is pass-by-reference by default
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type countsErr struct {
	counts map[string]int
	err    error
}

func main() {
	ccounts := make(chan countsErr)
	for _, p := range os.Args[1:] {
		go workerScan(p, ccounts)
	}

	total_counts := make(map[string]int)

	for range os.Args[1:] {
		ce := <-ccounts
		if ce.err != nil {
			log.Fatal(ce.err)
		}

		for w, c := range ce.counts {
			total_counts[w] += c
		}
	}

	for w, c := range total_counts {
		fmt.Printf("word: %s | count: %d\n", w, c)
	}
}

func workerScan(filename string, ccounts chan countsErr) {
	counts := make(map[string]int)
	err := scan(filename, counts)
	ccounts <- countsErr{counts, err}
}

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