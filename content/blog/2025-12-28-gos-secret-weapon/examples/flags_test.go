package examples

import (
	"fmt"
)

// ExampleCSVFlag demonstrates implementing flag.Value
func ExampleCSVFlag() {
	var tags CSVFlag
	tags.Set("production,api")
	fmt.Println(tags.String())
	// Output: production,api
}
