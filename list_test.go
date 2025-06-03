package list_test

import (
	"fmt"

	"github.com/eccles/go-list"
)

// Head tests.

func ExampleHead_singleton() {
	a := []int{10, 20, 20, 20, 30}
	b := list.Head(a)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 1 [10]
}

func ExampleHead_duplicates() {
	a := []int{20, 20, 20, 30}
	b := list.Head(a)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 3 [20 20 20]
}

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

func ExampleConflate_noduplicates() {
	a := []int{10, 20, 30}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 3 [10 20 30]
}

func ExampleConflate_allduplicates() {
	a := []int{20, 20, 20}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 0 []
}

func ExampleConflate_allduplicates1() {
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
	a := []int{10, 20, 20, 20, 30}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 2 [10 30]
}

func ExampleConflate_midduplicates1() {
	a := []int{10, 20, 20, 30}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 4 [10 20 20 30]
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

func ExampleConflate_trailingduplicate1() {
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
	a := []int{10, 10, 10, 20, 30}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 2 [20 30]
}

func ExampleConflate_leadingduplicate1() {
	a := []int{10, 10, 20, 30}
	b := list.Conflate(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 4 [10 10 20 30]
}
