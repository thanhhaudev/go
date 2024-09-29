package main

/*
Go is a call by value language, so when we pass a slice to a function, we are passing a copy of the slice header, not the underlying array.
Passing a slice to the append function actually passes a copy of the slice header, not the underlying array.

Here is a visual representation of a slice header:
+---------+---------+---------+
| Pointer |  Length | Capacity|
+---------+---------+---------+

+ Changes to the slice header in the function are not reflected in the original slice. The original slice is not modified.
+ Changes to the underlying array in the function are reflected in the original slice. The original slice is modified. e.g. append function or modifying the elements of the slice directly.

How append works:
+ If the slice has enough capacity, append will add the new elements to the slice and return the original slice.
+ If the slice does not have enough capacity, append will create a new underlying array with double the capacity, copy the elements from the original array to the new array, add the new elements to the new array, and return the new slice.
*/
func main() {
	var x []int

	x = append(x, 1)    // append 1 element to the slice
	x = append(x, 2, 3) // append 2 elements to the slice

	y := []int{4, 5}

	x = append(x, y...) // append a slice to the slice using the variadic parameter syntax

	/*
		It's compile-time error:
		+ if we use the variadic parameter syntax to append a slice to a slice of a different type, for example: x = append(x, []string{"a", "b"}...)
		+ if we forget to assign the value returned by the append function to the slice, for example: append(x, 1)
		+ The function adds elements to the slice, but the slice header is a copy, so the original slice is not modified. If we want to modify the original slice, we need to set the returned slice back to the variable that holds the original slice.
	*/
}
