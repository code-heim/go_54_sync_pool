package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a sync.Pool with a New function to initialize objects.
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Allocating new []byte slice")
			return make([]byte, 1024) // Allocate a 1KB slice
		},
	}

	// Get a new object from the pool
	// The Pool allocates since the pool is empty.
	obj := pool.Get().([]byte)
	fmt.Printf("Got object from pool of length: %d\n", len(obj))

	// Put the object back into the pool.
	pool.Put(obj)

	// Get the object again
	// this time it is reused from the pool).
	reusedObj := pool.Get().([]byte)
	fmt.Printf("Got reused object from pool of length: %d\n",
		len(reusedObj))

	// Put the object back into the pool again.
	pool.Put(reusedObj)
}
