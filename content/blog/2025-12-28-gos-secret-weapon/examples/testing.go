package examples

import "testing"

// ============================================================================
// ACCEPTING INTERFACES
// ============================================================================

// assertEqual is a reusable test helper that accepts testing.TB.
func assertEqual(tb testing.TB, expected, actual any) {
	// Helper() ensures error messages point to the caller's line, not here.
	tb.Helper()
	if expected != actual {
		tb.Fatalf("expected %v, got %v", expected, actual)
	}
}
