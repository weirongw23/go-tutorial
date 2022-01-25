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

func select_statement() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2s"
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

// Worker Pool
// Queue of work to be done and 
// multiple concurrent workers pulling work
// off the queue
func fib(n int) int {
	if n <= 1 {
        return n
    }
    
    dp := make([]int, n + 1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}

	return dp[n]
}

func worker_job(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

// Main itself is a routine
func main() {
	jobs := make(chan int, 80)
	results := make(chan int, 80)

	for i := 0; i < 4; i++ {
		go worker_job(jobs, results)
	}

	for i := 0; i < 80; i++ {
		jobs <- i
	}

	close(jobs)

	for i := 0; i < 80; i++ {
		fmt.Println(<-results)
	}
}