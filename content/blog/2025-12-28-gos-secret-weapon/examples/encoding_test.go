package examples

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ExampleToText demonstrates accepting encoding.TextMarshaler
func ExampleToText() {
	status := Running
	text, err := ToText(status)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Status:", text)
	}
	// Output: Status: running
}

// ExampleStatus demonstrates implementing encoding.TextMarshaler
func ExampleStatus() {
	status := Running
	b, _ := json.Marshal(status)
	fmt.Println(string(b))
	// Output: "running"
}

// ExampleTimestamp demonstrates implementing json.Marshaler
func ExampleTimestamp() {
	ts := Timestamp(1735689600)
	b, _ := json.Marshal(ts)
	// Output includes timezone offset which varies by system
	fmt.Printf("Timestamp contains 2025: %v\n", strings.Contains(string(b), "2025"))
	// Output: Timestamp contains 2025: true
}
