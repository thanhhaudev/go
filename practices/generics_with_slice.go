package main

import (
	"fmt"
	"golang.org/x/exp/constraints" // import the constraints package
)

// findMaxInSlice returns the maximum value in a slice of ordered types.
// The type parameter T is constrained to the Ordered interface. This means that the type T has constraints that it must satisfy.
func findMaxInSlice[T constraints.Ordered](slice []T) T {
	if len(slice) == 0 {
		panic("empty slice")
	}

	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}

	return max
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	// find the maximum value in the slice of integers
	fmt.Println(findMaxInSlice(a)) // 5

	fmt.Println(findMaxInSlice(b)) // 5.5
}
