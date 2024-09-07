package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	allocationCnt := 0
	// Create a sync.Pool with a New function to initialize objects.
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Print(".")
			allocationCnt++
			return make([]byte, 1024) // Allocate a 1KB slice
		},
	}

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			obj := pool.Get().([]byte)
			fmt.Print("-")
			// Simulate some work
			time.Sleep(100 * time.Millisecond)
			pool.Put(obj)
			wg.Done()
		}()
		time.Sleep(10 * time.Millisecond)
	}

	wg.Wait()

	fmt.Println("\n Number of allocations: ", allocationCnt)
}
