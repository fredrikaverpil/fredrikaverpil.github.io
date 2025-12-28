package examples

import (
	"fmt"
)

// ExampleMax demonstrates cmp.Ordered constraint in generic functions
func ExampleMax() {
	fmt.Println(Max(10, 20))
	fmt.Println(Max(3.14, 2.71))
	fmt.Println(Max("apple", "banana"))
	// Output:
	// 20
	// 3.14
	// banana
}
