package list

import (
	"testing"
)

func BenchmarkHead(b *testing.B) {
	for b.Loop() {
		Head([]int{10, 20, 20, 20, 30})
	}
}

func BenchmarkConflate(b *testing.B) {
	for b.Loop() {
		Conflate([]int{10, 20, 20, 20, 30}, 2)
	}
}
