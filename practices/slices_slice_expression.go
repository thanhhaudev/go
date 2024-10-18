package main

import "fmt"

/*
- Slice expressions are used to create a new slice by specifying a lower and upper bound, separated by a colon.
- New slices created by slice expressions share the same underlying array as the original slice.
- The lower bound is the index of the first element of the new slice, and the upper bound is the index of the element after the last element of the new slice.
- The upper bound is optional; if omitted, the new slice goes to the end of the original slice.
- The lower bound is also optional; if omitted, the slice starts from the beginning of the original slice.
- The new slice has a length and capacity determined by the start and end indices of the slice expression.
*/

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := x[0:2] // this creates a new slice from the first element to the third element of the slice x, the length of y is 2 and the capacity is 5
	z := x[1:3] // this creates a new slice from the second element to the fourth element of the slice x, the length of z is 2 and the capacity is 4
	e := x[:]   // this creates a new slice from the first element to the last element of the slice x, the length of e is 5 and the capacity is 5

	/*
		In Go, the capacity of a slice created by a slice expression is determined by the capacity of the original slice minus the starting index of the new slice.

		y := x[0:2]: The capacity of y is the capacity of x (which is 5) minus the starting index (0). So, the capacity of y is 5 - 0 = 5.
		z := x[1:3]: The capacity of z is the capacity of x (which is 5) minus the starting index (1). So, the capacity of z is 5 - 1 = 4.
		This is why y has a capacity of 5 and z has a capacity of 4.
	*/
	fmt.Println(x, len(x), cap(x)) // [1 2 3 4 5] 5 5
	fmt.Println(y, len(y), cap(y)) // [1 2] 2 5
	fmt.Println(z, len(z), cap(z)) // [2 3] 2 4
	fmt.Println(e, len(e), cap(e)) // [1 2 3 4 5] 5 5

	// Remember that when we take a slice from another slice, the new slice shares the same underlying array as the original slice. This means that changes to the new slice will affect the original slice
	// this called overlapping slices
	x = []int{1, 2, 3, 4}
	y = x[:2] // y is [1 2]
	z = x[1:] // z is [2 3 4]

	x[1] = 20            // this will change 1 to 20 in the original slice x
	y[0] = 10            // this will change 1 to 10 in the new slice y
	z[1] = 30            // this will change 3 to 30 in the new slice z
	fmt.Println("x:", x) // [10 20 30 4]
	fmt.Println("y:", y) // [10 20]
	fmt.Println("z:", z) // [20 30 4]

	// append() makes overlapping slices more complicated:
	x = []int{1, 2, 3, 4}
	y = x[:2]                                              // y is [1 2]
	fmt.Printf("cap(x): %d, cap(y): %d\n", cap(x), cap(y)) // cap(x): 4, cap(y): 4
	y = append(y, 99)                                      // y changes to [1 2 99]
	fmt.Println("x:", x)                                   // [1 2 99 4] <- the original slice x also changes
	fmt.Println("y:", y)                                   // [1 2 99]

	// Refer to line 21 for the explanation of the capacity of y.
	// In this case, any unused capacity in the original slice is also shared with the new slice.
	// When we create y from x, the length is set to 2 but the capacity is 4, same as x. (Refer to `slices_append_capacity.go` to understand how the capacity is calculated.)
	// So when we append 99 to y, it will change the original slice's 3rd element to 99 because the available capacity is shared between x and y (they have the same memory address).

	// More complicated example:
	x = make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)                                                  // [1 2 3 4]
	y = x[:2]                                                                  // [1 2]
	z = x[2:]                                                                  // [3 4]
	fmt.Printf("cap(x): %d, cap(y): %d, cap(z): %d\n", cap(x), cap(y), cap(z)) // cap(x): 5, cap(y): 5, cap(z): 3
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println("x:", x) // [1 2 30 40 70]
	fmt.Println("y:", y) // [1 2 30 40 70]
	fmt.Println("z:", z) // [30 40 70]

	// To avoid complicated slice situations, we should never use append() on a slice that is created from another slice.
	// Instead, we can use `full slice expressions` to create a new slice with a specified length and capacity.
	// A `full slice expression` includes a third index which indicates the maximum capacity of the new slice.
	// The capacity of the new slice is determined by subtracting the starting index from the third index. This means we can specify the capacity of the new slice.
	// If we provide wrong capacity, the compiler will throw runtime error: `panic: runtime error: slice bounds out of range [:3] with capacity 2`
	// This way, the new slice will have its own underlying array and will not affect the original slice when we append() elements to it.

	x = make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)                                                  // [1 2 3 4]
	y = x[:2:2]                                                                // y is [1 2] with a capacity of 2 (2-0)
	z = x[2:4:4]                                                               // z is [3 4] with a capacity of 2 (4-2)
	fmt.Printf("cap(x): %d, cap(y): %d, cap(z): %d\n", cap(x), cap(y), cap(z)) // cap(x): 5, cap(y): 2, cap(z): 2

	// these loop will print the address of each element in the slice, they will have the same address

	for i := range x {
		fmt.Printf("x: element %d address %p\n", i, &x[i])
	}

	for i := range y {
		fmt.Printf("y: element %d address %p\n", i, &y[i])
	}

	for i := range z {
		fmt.Printf("z: element %d address %p\n", i, &z[i])
	}

	y = append(y, 99)
	z = append(z, 100)
	x = append(x, 200)
	fmt.Println("x:", x) // [1 2 3 4 200]
	fmt.Println("y:", y) // [1 2 99]
	fmt.Println("z:", z) // [3 4 100]

	// After appending items to y and z, they will have different addresses because when the new length exceeds the capacity,
	// a new underlying array is created and the elements are copied to the new array.
	// So the new slice will have its own underlying array and will not affect the original slice.
}
