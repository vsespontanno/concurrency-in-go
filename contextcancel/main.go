package main

import (
	"context"
	"fmt"
	"time"
)

func longOperation(ctx context.Context) error {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Printf("Working... step %d\n", i)
			time.Sleep(500 * time.Millisecond)
		}
	}

	fmt.Println("Operation completed successfully")
	return nil
}

func main() {
	fmt.Println("=== Scenario 1: Operation completes successfully ===")
	ctx1, cancel1 := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel1()

	err := longOperation(ctx1)
	if err != nil {
		fmt.Printf("Result: Failed with error: %v\n", err)
	} else {
		fmt.Println("Result: Success!")
	}

	fmt.Println("\n=== Scenario 2: Operation times out ===")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel2()

	err = longOperation(ctx2)
	if err != nil {
		fmt.Printf("Result: Failed with error: %v\n", err)
	} else {
		fmt.Println("Result: Success!")
	}

	fmt.Println("\n=== Scenario 3: Manual cancellation ===")
	ctx3, cancel3 := context.WithCancel(context.Background())

	time.Sleep(1 * time.Second)
	fmt.Println("Manually cancelling operation...")
	cancel3()

	err = longOperation(ctx3)
	if err != nil {
		fmt.Printf("Result: Failed with error: %v\n", err)
	} else {
		fmt.Println("Result: Success!")
	}
}
