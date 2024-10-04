package main

import "fmt"

/*
Instead of declaring a slice using a slice literal or nil, we can use the built-in `make` function to create a slice with a specified length and capacity.
If you do not provide the capacity when using the make function to create a slice, the capacity will be equal to the length.
*/

func main() {
	x := make([]int, 5)            // this creates an empty slice with a length and capacity of 5
	fmt.Println(x, len(x), cap(x)) // [0 0 0 0 0] 5 5
	fmt.Println(x[0])              // 0 because this is the default value for an int

	y := make([]int, 5, 10)        // this creates an empty slice with a length of 5 and a capacity of 10
	fmt.Println(y, len(y), cap(y)) // [0 0 0 0 0] 5 10

	z := make([]int, 0, 10)        // this creates an empty slice with a length of 0 and a capacity of 10
	fmt.Println(z, len(z), cap(z)) // [] 0 10
	// In this case, the slice we have non-nil, but it has a length of 0. This is useful when we want to append elements to the slice later without allocating more memory.
	// If we append elements to the slice like this: z = append(z, 1, 2, 3, 4, 5), the slice will have a length of 5 and a capacity of 10.
	// fmt.Println(z, len(z), cap(z)) // [1 2 3 4 5] 5 10

	// One common beginner mistake is to try to populate those initial elements using append()
	// x = append(x, 1, 2, 3, 4, 5) // this will not work because the slice has a length of 5 and a capacity of 5
	// fmt.Println(x, len(x), cap(x)) // [0 0 0 0 0 1 2 3 4 5] 10 10
}
