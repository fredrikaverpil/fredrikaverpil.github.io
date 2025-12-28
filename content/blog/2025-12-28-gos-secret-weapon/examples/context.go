package examples

import (
	"context"
	"time"
)

// ============================================================================
// ACCEPTING INTERFACES
// ============================================================================

// GetUser demonstrates accepting context.Context for cancellation and deadline support.
func GetUser(ctx context.Context, id int) (*User, error) {
	// Simulate work that respects context cancellation
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(100 * time.Millisecond):
		return &User{Name: "Alice"}, nil
	}
}

// ============================================================================
// TYPES
// ============================================================================

// User represents a user from database.
type User struct {
	Name string
}