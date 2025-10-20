package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	sem     chan struct{}
	max     int
	maxTime time.Duration
}

func New(max int, maxTime time.Duration) *RateLimiter {
	sem := make(chan struct{}, max)
	rl := &RateLimiter{
		sem:     sem,
		max:     max,
		maxTime: maxTime,
	}
	for i := 0; i < max; i++ {
		rl.sem <- struct{}{}
	}

	go func() {
		ticker := time.NewTicker(maxTime / time.Duration(max))
		defer ticker.Stop()

		for range ticker.C {
			select {
			case rl.sem <- struct{}{}:
			default:
			}
		}

	}()
	return rl
}

func (r *RateLimiter) Proceed(fn func()) {
	<-r.sem
	fn()
}

func main() {
	rl := New(2, time.Second)
	var wg sync.WaitGroup

	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			rl.Proceed(func() {
				fmt.Printf("API call %d at %v\n", i, time.Now())
				time.Sleep(100 * time.Millisecond)
			})
		}(i)
	}

	wg.Wait()
}
