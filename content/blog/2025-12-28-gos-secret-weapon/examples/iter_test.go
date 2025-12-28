package examples

import (
	"fmt"
	"slices"
)

// ExampleCollect demonstrates collecting iter.Seq into a slice.
func ExampleCollect() {
	nums := Collect(slices.Values([]int{1, 2, 3}))
	fmt.Println(nums)
	// Output:
	// [1 2 3]
}

// ExampleFilter demonstrates filtering a sequence with a predicate.
func ExampleFilter() {
	nums := Collect(Filter(slices.Values([]int{1, 2, 3, 4}), func(n int) bool {
		return n%2 == 0
	}))
	fmt.Println(nums)
	// Output:
	// [2 4]
}

// ExampleStack demonstrates that custom types can implement iter.Seq.
func ExampleStack() {
	stack := &Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	nums := Collect(stack.All())
	fmt.Println(nums)
	// Output:
	// [3 2 1]
}
