package examples

import "cmp"

// ============================================================================
// ACCEPTING INTERFACES
// ============================================================================

// Max returns the maximum of two ordered values using cmp.Ordered constraint.
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
