package examples

import (
	"errors"
	"fmt"
)

// ExampleAnnotate demonstrates wrapping an error with additional context.
func ExampleAnnotate() {
	err := fmt.Errorf("file not found")
	err = Annotate(err, "loading config")
	fmt.Println(err)
	// Output: loading config: file not found
}

// ExampleToJSON demonstrates marshaling any value to JSON.
func ExampleToJSON() {
	fmt.Println(ToJSON(42))
	fmt.Println(ToJSON(map[string]int{"a": 1}))
	// Output:
	// 42
	// {"a":1}
}

// ExampleFind demonstrates searching for a target in different types.
func ExampleFind() {
	fmt.Println(Find([]int{1, 2, 3}, 2))
	fmt.Println(Find([]string{"a", "b"}, "b"))
	fmt.Println(Find([]int{1, 2, 3}, 99))
	// Output:
	// 1
	// 1
	// -1
}

// ExampleValidationError demonstrates a custom error type.
func ExampleValidationError() {
	err := &ValidationError{Field: "email", Message: "invalid format"}
	fmt.Println(err)
	// Output: email: invalid format
}

// ExampleQueryError demonstrates error chaining.
func ExampleQueryError() {
	inner := errors.New("no rows")
	err := &QueryError{Query: "SELECT *", Err: inner}
	fmt.Println(err)
	fmt.Println(errors.Unwrap(err))
	// Output:
	// SELECT *: no rows
	// no rows
}
