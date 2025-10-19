package main

import (
	"fmt"
	"sync"
)

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)
	ch1 <- 1
	ch1 <- 5
	ch2 <- 1
	ch2 <- 5
	ch2 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	close(ch1)
	close(ch2)
	close(ch3)
	for v := range fanIn(ch1, ch2, ch3) {
		fmt.Print(v)
	}
}
