package examples

import "iter"

// ============================================================================
// ACCEPTING INTERFACES
// ============================================================================

// Collect gathers any sequence into a slice, accepting iter.Seq.
// Works with slices.Values(), maps.Keys(), custom iterators, etc.
func Collect[T any](seq iter.Seq[T]) []T {
	var result []T
	for v := range seq {
		result = append(result, v)
	}
	return result
}

// Filter accepts iter.Seq and returns a filtered sequence with a predicate.
func Filter[T any](seq iter.Seq[T], keep func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if keep(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// ============================================================================
// IMPLEMENTING INTERFACES
// ============================================================================

// Stack is a LIFO data structure that implements iter.Seq for iteration.
type Stack[T any] struct {
	items []T
}

// Push adds an item to the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// All returns an iterator over the stack in LIFO order, implementing iter.Seq.
// This allows the stack to work with for-range and all iter functions.
func (s *Stack[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(s.items) - 1; i >= 0; i-- {
			if !yield(s.items[i]) {
				return
			}
		}
	}
}
