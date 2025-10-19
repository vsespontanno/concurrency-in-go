package main

import (
	"fmt"
	"time"
)

func waitWithTimeout(d time.Duration, ch <-chan int) (int, error) {
	select {
	case v := <-ch:
		return v, nil
	case <-time.After(d):
		return 0, fmt.Errorf("timeout: %d seconds", d)
	}
}
