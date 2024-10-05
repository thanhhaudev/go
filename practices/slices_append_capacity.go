package main

import "fmt"

/*
- Every slice has a capacity, which is the maximum number of elements the slice can hold without allocating more memory. This can be larger than the length of the slice.
- When we append elements to a slice, one or more values is added to the end of the slice.
- If the slice has enough capacity to hold the new elements, the elements are added to the slice in place.
- If we try to add more elements when the length equals (or exceeds) the capacity, a new underlying array is created with double the capacity, the elements from the original array are copied to the new array, and the new elements are added to end, and the slice header is updated to point to the new array.
- The rules as of Go 1.14 are double the capacity if the length is less than 1024, and 25% more capacity if the length is 1024 or more.
- We can use the built-in `cap` function to get the capacity of a slice.
*/
func main() {
	var x []int
	fmt.Println(x, len(x), cap(x)) // 0 0
	x = append(x, 1)
	fmt.Println(x, len(x), cap(x)) // [1] 1 1
	x = append(x, 2)
	fmt.Println(x, len(x), cap(x)) // [1 2] 2 2
	x = append(x, 3)
	fmt.Println(x, len(x), cap(x)) // [1 2 3] 3 4
	x = append(x, 4)
	fmt.Println(x, len(x), cap(x)) // [1 2 3 4] 4 4
	x = append(x, 5)
	fmt.Println(x, len(x), cap(x)) // [1 2 3 4 5] 5 8
	x = append(x, 6)
	fmt.Println(x, len(x), cap(x)) // [1 2 3 4 5 6] 6 8
	x = append(x, 7)
	fmt.Println(x, len(x), cap(x)) // [1 2 3 4 5 6 7] 7 8
	x = append(x, 8)
	fmt.Println(x, len(x), cap(x)) // [1 2 3 4 5 6 7 8] 8 8
	x = append(x, 9)
	fmt.Println(x, len(x), cap(x)) // [1 2 3 4 5 6 7 8 9] 9 16

	/*
		The primary goal is to minimize the number of times the slice needs to be grown. This is because growing a slice requires allocating a new array and copying the elements from the old array to the new array.
		- If we know the maximum number of elements the slice will hold, we can use the built-in `make` function to create a slice with a specified length and capacity.
		- If we do not provide the capacity when using the make function to create a slice, the capacity will be equal to the length.
	*/
}
