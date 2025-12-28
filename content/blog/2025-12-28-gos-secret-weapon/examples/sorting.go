package examples

import "sort"

// ============================================================================
// ACCEPTING INTERFACES
// ============================================================================

// IsSorted checks if any sort.Interface is sorted.
func IsSorted(data sort.Interface) bool {
	n := data.Len()
	for i := 1; i < n; i++ {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

// ============================================================================
// IMPLEMENTING INTERFACES
// ============================================================================

// Person is a struct for sorting examples.
type Person struct {
	Name string
	Age  int
}

// ByAge sorts people by age and implements sort.Interface.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// IntHeap implements heap.Interface for use with container/heap.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
