package examples

import "sync"

// ============================================================================
// ACCEPTING INTERFACES
// ============================================================================

// WithLock safely executes a function with a lock using the sync.Locker interface.
func WithLock(l sync.Locker, fn func()) {
	l.Lock()
	defer l.Unlock()
	fn()
}
