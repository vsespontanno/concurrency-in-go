package main

import (
	"sync"
)

func worker(WID int, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		println("worker", WID, "started  job", j)
		result <- j * 2
	}
}

func main() {
	numberOfWorkers := 3
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var workersWG sync.WaitGroup

	// Запускаем воркеров
	for w := 1; w <= numberOfWorkers; w++ {
		workersWG.Add(1)
		go worker(w, jobs, results, &workersWG)
	}

	go func() {
		for j := 1; j <= 100; j++ {
			jobs <- j
		}
		close(jobs)
	}()
	for a := 1; a <= 100; a++ {
		result := <-results
		println("result", result)
	}

	workersWG.Wait()

}
