package examples

import (
	"container/heap"
	"fmt"
)

// ExampleIsSorted demonstrates accepting sort.Interface
func ExampleIsSorted() {
	people := ByAge{{Name: "Alice", Age: 20}, {Name: "Bob", Age: 25}}
	fmt.Println("Is sorted:", IsSorted(people))
	// Output: Is sorted: true
}

// ExampleIntHeap demonstrates implementing heap.Interface
func ExampleIntHeap() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	min := heap.Pop(h)
	fmt.Println("Min:", min)
	// Output: Min: 1
}
