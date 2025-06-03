package list

import (
	"testing"
)

func BenchmarkHead_int(b *testing.B) {
	for b.Loop() {
		Head([]int{10, 20, 20, 20, 30})
	}
}

func BenchmarkHead_string(b *testing.B) {
	for b.Loop() {
		Head([]string{"10", "20", "20", "20", "30"})
	}
}

func BenchmarkConflate_int(b *testing.B) {
	for b.Loop() {
		Conflate([]int{10, 20, 20, 20, 30}, 2)
	}
}

func BenchmarkConflate_string(b *testing.B) {
	for b.Loop() {
		Conflate([]string{"10", "20", "20", "20", "30"}, 2)
	}
}
