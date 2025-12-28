package examples

import (
	"encoding/json"
	"fmt"
)

// ============================================================================
// ACCEPTING INTERFACES
// ============================================================================

// Annotate wraps any error with additional context using error composition.
func Annotate(err error, op string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", op, err)
}

// ToJSON marshals any value to JSON using the empty interface.
func ToJSON(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// Find searches for a target in a slice using generics.
func Find[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// ============================================================================
// IMPLEMENTING INTERFACES
// ============================================================================

// ValidationError is a custom error type that implements the error interface.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// QueryError is an error type that chains errors using the Unwrap method.
type QueryError struct {
	Query string
	Err   error
}

func (e *QueryError) Error() string {
	return e.Query + ": " + e.Err.Error()
}

func (e *QueryError) Unwrap() error {
	return e.Err
}
