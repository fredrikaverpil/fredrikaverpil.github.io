package examples

import "fmt"

// ExampleFormatLabel demonstrates accepting fmt.Stringer.
func ExampleFormatLabel() {
	id := UserID(42)
	result := FormatLabel(id)
	fmt.Println(result)
	// Output: [user-42]
}

// ExampleUserID demonstrates implementing fmt.Stringer.
func ExampleUserID() {
	id := UserID(42)
	fmt.Println(id)
	// Output: user-42
}

// ExampleConfig demonstrates implementing fmt.GoStringer for redaction.
func ExampleConfig() {
	cfg := Config{Host: "localhost", Password: "secret"}
	fmt.Printf("%#v\n", cfg)
	// Output: Config{Host:"localhost", Password:"***"}
}
