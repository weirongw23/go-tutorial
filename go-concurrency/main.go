package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Wait Groups and Semaphores
// Channel Syntax
// <-chan: read-only (receive)
// chan<-: write (send, can close)
func worker(id int, wg *sync.WaitGroup, sem <-chan bool) {
	// Defer executes like a stack - LIFO
	defer wg.Done()
	defer log.Printf("#%d done", id)

	log.Printf("#%d starting", id)
	time.Sleep(time.Second)
	<-sem
}

func wait_group() {
	// Ensures we have at most 5 workers executing at once
	// Useful for something like a memory-heavy operation
	var wg sync.WaitGroup
	sem := make(chan bool, 5)
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		sem <- true
		go worker(i, &wg, sem)
	}

	wg.Wait()
	log.Printf("all done")
}

// Count function
func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func wait_group_intro() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		count("sheep")
		defer wg.Done()
	}()

	wg.Wait()
}

func count_channels(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c) // close channel on sending side
}

func channel_intro_1() {
	c := make(chan string)
	go count_channels("sheep", c)

	for {
		msg, open := <- c

		if !open {
			break
		}

		fmt.Println(msg)
	}
}

func channel_intro_2() {
	c := make(chan string)
	go count_channels("sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

// Main itself is a routine
func main() {
	channel_intro_1()
}