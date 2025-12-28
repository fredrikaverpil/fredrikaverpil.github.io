package examples

import (
	"fmt"
	"sync"
)

// ExampleWithLock demonstrates sync.Locker usage
func ExampleWithLock() {
	var mu sync.Mutex
	var counter int

	WithLock(&mu, func() {
		counter++
	})

	fmt.Println("Counter:", counter)
	// Output: Counter: 1
}
