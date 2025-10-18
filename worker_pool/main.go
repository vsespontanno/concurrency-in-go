package main

import (
	"fmt"
	"sync"
)

func main() {
	numberOfWorkers := 2
	numberOfJobs := 5
	jobs := make(chan int, numberOfJobs)
	results := make(chan int, numberOfJobs)
	var wg sync.WaitGroup
	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		i := i
		go worker(i, jobs, results, &wg)

	}
	for i := 0; i < numberOfJobs; i++ {
		jobs <- i
	}
	close(jobs)
	for a := 0; a < numberOfJobs; a++ {
		<-results
	}

	wg.Wait()
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		results <- j * 2
		fmt.Println("worker", id, "processed job", j)
	}
}
