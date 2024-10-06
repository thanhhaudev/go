package main

import "fmt"

/*
- Slice expressions are used to create a new slice by specifying a lower and upper bound, separated by a colon.
- New slices created by slice expressions share the same underlying array as the original slice.
- The lower bound is the index of the first element of the slice, and the upper bound is the index of the element after the last element of the slice.
- The upper bound is optional; if omitted, the slice goes to the end of the original slice.
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
}
