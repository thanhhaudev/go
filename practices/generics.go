package main

import "fmt"

/*
- Generics in Go is a feature that allows us to write functions, data structures, and algorithms that work with any type.
This is a powerful feature that can help us write more flexible and reusable code.
*/

// Print this function can accept any type
// The type of the parameter is specified using the type parameter T
// The type parameter T is defined using the syntax [T any]
// The type parameter T is used as the type of the parameter s
// any is a type constraint that allows the function to accept any type (alias for interface{})
func Print[T any](s T) {
	fmt.Printf("%T\n", s)
}

// Swap this function can accept two different types
func Swap[T1, T2 any](a T1, b T2) (T2, T1) {
	return b, a
}

// Box this is a generic type
type Box[T any] struct {
	value T
}

// GetValue this is a generic method that returns the value of the Box
func (b Box[T]) GetValue() T {
	return b.value
}

func main() {
	// generic type with function
	Print([]int{1, 2, 3})                   // []int
	Print([]string{"a", "b", "c"})          // []string
	Print([]float64{1.1, 2.2, 3.3})         // []float64
	Print([]bool{true, false, true})        // []bool
	Print([]interface{}{1, "a", 1.1, true}) // []interface {}
	Print("hello")                          // string

	// generic type with struct
	boxInt := Box[int]{value: 10}
	fmt.Printf("%T %v\n", boxInt.GetValue(), boxInt.GetValue()) // int 10

	boxString := Box[string]{value: "hello"}
	fmt.Printf("%T %v\n", boxString.GetValue(), boxString.GetValue()) // string hello

	// multiple type parameters
	a, b := Swap(1, "hello")
	fmt.Println(a, b) // hello 1
}
