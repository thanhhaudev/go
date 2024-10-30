package main

import "fmt"

// if we need to create a slice that doesn't share the same underlying array with the original slice, we can use the copy function

func main() {
	src := []int{1, 2, 3, 4}
	dst := make([]int, 4)
	num := copy(dst, src) // if we don't need to know the number of elements copied, we can ignore the return value by don't assign it to any variable

	// copy(dst, src) copies elements from the src slice to the dst slice.
	// It returns the number of elements actually copied, which is the smaller length
	// of either src or dst. If dst is shorter than src, only len(dst) elements are copied.
	// If src is shorter than dst, only len(src) elements are copied.
	// Capacity of dst and src doesn't matter, only the length of the slices matters.
	//
	// The copy function does not share the same underlying array with the original slice. It creates a new array and copies the elements from the original slice to the new array.
	// This means that changes to the new slice will not affect the original slice.
	// The copy function is useful when we need to create a new slice that does not share the same underlying array with the original slice.
	// The copy function is also useful when we need to copy the elements of a slice to another slice.

	fmt.Printf("src: %v, dst: %v, num: %d\n", src, dst, num) // src: [1 2 3 4], dst: [1 2 3 4], num: 4

	// If we want to copy only a subset of the elements from the original slice to the new slice, we can use the slice expression to specify the range of elements to copy.
	dst = make([]int, 2)
	copy(dst, src[1:3])                        // copy elements from index 1 to 2 (3-1) from src to dst
	fmt.Printf("src: %v, dst: %v\n", src, dst) // src: [1 2 3 4], dst: [2 3], num: 2

	// The copy function allows us to copy elements between two slices, even if those slices overlap in the same underlying array.
	x := []int{1, 2, 3, 4}
	num = copy(x[1:], x)                   // copy elements from x into x[1:], starting from index 1
	fmt.Printf("x: %v, num: %d\n", x, num) // x: [1 1 2 3], num: 3
	// We are copying elements from the original slice x into a slice that overlaps with it, x[1:].
	// Initially, x[1:] is [2, 3, 4]. Since x[1:] has length 3, only the first 3 elements from x ([1, 2, 3]) are copied into x[1:].
	// This results in x being modified to [1, 1, 2, 3].

	// We also can use copy with arrays by taking a slice of the array. An arrays can be src or dst.
	arr := [4]int{1, 2, 3, 4}
	a1 := make([]int, 4)
	copy(a1, arr[:])
	fmt.Printf("a1: %v\n", a1) // a1 [1 2 3 4]
	copy(arr[:], a1)
	fmt.Printf("arr: %v\n", arr) // arr [1 2 3 4]
}
