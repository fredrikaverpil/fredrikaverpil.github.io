package examples

import "testing"

func TestAssertEqual(t *testing.T) {
	assertEqual(t, 1, 1)
}

func BenchmarkAssertEqual(b *testing.B) {
	for b.Loop() {
		assertEqual(b, 1, 1)
	}
}
