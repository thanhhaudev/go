package main

import "fmt"

/*
arrays in Go are rarely used because they have a fixed size, which makes them less flexible than slices.
Go considers the size of an array to be part of its type, so [2]int and [3]int are different types.
We can't use a variable to specify the size of an array, so we can't create an array with a size that is determined at runtime.
Only use arrays when we know exactly how many elements we need ahead of time.
Because size is part of the type, we can't write a function that accepts an array of any size.
*/
func main() {
	var a = [2]int{1, 2} // create an array with 2 elements
	var b [3]int         // this mean b is an array with 3 elements, but the elements are not initialized yet, so the default value of each element is 0 (default value of int)
	var c = [10]int{
		1, 5: 99, 9: 100,
	} // this called sparse array, the first element is 1, the fifth element is 99, the ninth element is 100, the rest of the elements are 0
	var d = [...]int{
		1, 2, 3,
	} // when using an array literal, we can use ... to let the compiler count the number of elements for us

	fmt.Println(a) // [1 2]
	fmt.Println(b) // [0 0 0]
	fmt.Println(c) // [1 0 0 0 99 0 0 0 0 100]
	fmt.Println(d) // [1 2 3]

	// we cant use == and != to compare two arrays to check if they are equal or not, we need to compare each element of the arrays
	// the equality operator == and the inequality operator != require the operands to be of the same type and value, position, and length
	// fmt.Println(a == b) // invalid operation: a == b (mismatched types [2]int and [3]int)
	// fmt.Println(a != b) // invalid operation: a != b (mismatched types [2]int and [3]int)
	var e = [2]int{1, 2}
	var f = [2]int{1, 2}
	var g = [2]int{2, 1}
	fmt.Println(e == f) // true
	fmt.Println(e == g) // false

	// Go only supports one-dimensional arrays, but we can create a multi-dimensional array by creating an array of arrays
	var h = [2][2]int{
		{1, 2},
		{3, 4},
	}

	fmt.Println(h) // [[1 2] [3 4]]

	// we can't read or write an element of an array that is out of bounds or use a negative index, it will cause a runtime error
	// fmt.Println(a[2]) // panic: runtime error: index out of range [2] with length 2
	// fmt.Println(a[-1]) // panic: runtime error: index out of range [-1] with length 2

	// we can use the len function to get the length of an array
	fmt.Println(len(a)) // 2

	// we can't use a type conversion to convert an array to another array with a different size.
	// fmt.Println([3]int(a)) // cannot convert a (type [2]int) to type [3]int
	// if two arrays have the same type, we can convert one array to another array with the same type
	// fmt.Println([2]int(a)) // [1 2]
}
