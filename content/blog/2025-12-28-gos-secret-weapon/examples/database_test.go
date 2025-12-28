package examples

import (
	"fmt"
)

// ExampleDebugSQLValue demonstrates accepting driver.Valuer
func ExampleDebugSQLValue() {
	data := JSONB{"test": true}
	result := DebugSQLValue(data)
	fmt.Println(result)
	// Output: SQL Value: {"test":true}
}

// ExampleJSONB demonstrates implementing sql.Scanner and driver.Valuer
func ExampleJSONB() {
	data := JSONB{"theme": "dark"}
	b, _ := data.Value()
	var data2 JSONB
	data2.Scan(b)
	fmt.Println("Theme:", data2["theme"])
	// Output: Theme: dark
}
