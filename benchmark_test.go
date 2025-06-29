package list

import (
	"testing"
)

// b.Loop is used but there may be some drawbacks:
//
// https://github.com/golang/go/issues/73137

func BenchmarkConflate_int_leadingduplicates(b *testing.B) {
	for b.Loop() {
		Conflate([]int{20, 20, 20, 30, 40, 50, 60}, 2)
	}
}

func BenchmarkConflate_int_midduplicates(b *testing.B) {
	for b.Loop() {
		Conflate([]int{10, 20, 20, 20, 30, 40, 50, 60}, 2)
	}
}

func BenchmarkConflate_int_trailingduplicates(b *testing.B) {
	for b.Loop() {
		Conflate([]int{10, 40, 50, 60, 20, 20, 20}, 2)
	}
}

func BenchmarkConflate_string_leadingduplicates(b *testing.B) {
	for b.Loop() {
		Conflate([]string{"20", "20", "20", "30", "40", "50", "60"}, 2)
	}
}

func BenchmarkConflate_string_midduplicates(b *testing.B) {
	for b.Loop() {
		Conflate([]string{"10", "20", "20", "20", "30", "40", "50", "60"}, 2)
	}
}

func BenchmarkConflate_string_trailingduplicates(b *testing.B) {
	for b.Loop() {
		Conflate([]string{"10", "40", "50", "60", "20", "20", "20"}, 2)
	}
}
