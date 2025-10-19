package main

import "sync"

type SafeCounter struct {
	iter int
	mux  *sync.Mutex
}

func New() *SafeCounter {
	iter := 0
	mux := &sync.Mutex{}
	return &SafeCounter{iter: iter, mux: mux}
}

func (c *SafeCounter) Inc() {
	c.mux.Lock()
	c.iter++
	c.mux.Unlock()
}

func (c *SafeCounter) Value() int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.iter
}

func main() {
	sc := New()
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sc.Inc()
		}()
	}
	wg.Wait()
	println(sc.Value(), sc.Value() == 1000)
}
