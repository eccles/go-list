package list_test

import (
	"fmt"

	"github.com/eccles/go-list"
)

// Conflate tests.

func ExampleConflate_empty() {
	a := []int{}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 0 []
}

func ExampleConflate_emptystring() {
	a := []string{}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 0 []
}

func ExampleConflate_noduplicates() {
	a := []int{10, 20, 30, 40, 50, 60}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 6 [10 20 30 40 50 60]
}

func ExampleConflate_noduplicatesstring() {
	a := []string{"10", "20", "30", "40", "50", "60"}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 6 [10 20 30 40 50 60]
}

func ExampleConflate_allduplicates() {
	a := []int{20, 20, 20, 20, 20, 20}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 0 []
}

func ExampleConflate_lessduplicates() {
	a := []int{20, 20}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 2 [20 20]
}

func ExampleConflate_midduplicates() {
	a := []int{10, 20, 20, 20, 30, 40, 50, 60}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 5 [10 30 40 50 60]
}

func ExampleConflate_lessmidduplicates() {
	a := []int{10, 20, 20, 30, 40, 50, 60}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 7 [10 20 20 30 40 50 60]
}

func ExampleConflate_trailingduplicate() {
	a := []int{10, 20, 30, 30, 30}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 2 [10 20]
}

func ExampleConflate_lesstrailingduplicate() {
	a := []int{10, 20, 30, 30}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 4 [10 20 30 30]
}

func ExampleConflate_leadingduplicate() {
	a := []int{10, 10, 10, 20, 30, 40, 50, 60}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 5 [20 30 40 50 60]
}

func ExampleConflate_lessleadingduplicate() {
	a := []int{10, 10, 20, 30}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 4 [10 10 20 30]
}
