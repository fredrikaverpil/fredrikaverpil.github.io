package examples

import (
	"context"
	"fmt"
	"time"
)

// ExampleGetUser demonstrates accepting context.Context for cancellation.
func ExampleGetUser() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	user, err := GetUser(ctx, 1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User:", user.Name)
	}
	// Output: User: Alice
}

// ExampleGetUser_timeout demonstrates context cancellation via timeout.
func ExampleGetUser_timeout() {
	// Set a timeout shorter than the simulated 100ms work in GetUser
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	_, err := GetUser(ctx, 1)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Output: Error: context deadline exceeded
}