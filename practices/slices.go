package main

import "fmt"

/*
slices in Go are more flexible than arrays because they have a dynamic size, which means we can change the size of a slice at runtime.
*/
func main() {
	var a = []int{1, 2} // create a slice with 2 elements
	var b = []int{
		1, 4: 99, 9: 100,
	} // this called sparse slice, just like arrays, we can specify only some elements of a slice, the rest of the elements are 0

	fmt.Println(a) // [1 2]
	fmt.Println(b) // [1 0 0 0 99 0 0 0 0 100]

	// we also can simulate multi-dimensional slices by creating a slice of slices
	var c = [][]int{
		{1, 2},
		{3, 4},
	}

	fmt.Println(c) // [[1 2] [3 4]]

	// just like arrays, we can't read or write an element of a slice that is out of bounds or use a negative index, it will cause a runtime error
	// fmt.Println(a[2]) // panic: runtime error: index out of range [2] with length 2
	// fmt.Println(a[-1]) // panic: runtime error: index out of range [-1] with length 2

	/* Different ways to create a slice */
	var d []int // create a slice with 0 elements
	// since no value is specified, the default value of a slice is nil
	// in Go, nil is an identifier that represents the lack of a value, it is the zero value for pointers, interfaces, channels, maps, slices, and functions
	// nil has no type, so we can compare it with any pointer, interface, channel, map, slice, or function value
	// a nil slice contains nothing, so its length and capacity are 0
	fmt.Println(d) // []

	// a slice isn't comparable, so we can't use == and != to compare two slices to check if they are equal or not
	// the only thing we can compare is nil
	fmt.Println(d == nil) // true
	// we can use reflect.DeepEqual to compare two slices

	// Different ways to create a slice
	// var data []int // this creates a nil slice (no memory is allocated)
	// data := []int{} // this creates an empty slice, and it is not nil (it has allocated memory, false when compared with nil)
	// data := make([]int, 0) // this creates an empty slice, and it is not nil (it has allocated memory)
}
