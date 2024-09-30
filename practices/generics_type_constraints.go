package main

import "fmt"

/*
- Generics type constraints allow us to specify the types that can be used with a generic function or type.
*/

type Adder interface {
	~int | ~float64
	// using `~` operator to allow the constraint to include multiple types that have the same underlying type as int or float64
	// this means it includes int, float64, and any custom types that are defined with int or float64 as their underlying type
	// if we want to allow only int and float64, just remove `~` operator
}

func Add[T Adder](a, b T) T {
	return a + b
}

type customInt int

func main() {
	fmt.Println(Add(1, 2))     // 3
	fmt.Println(Add(1.1, 2.2)) // 3.3

	// custom type with underlying type int
	fmt.Println(Add(customInt(1), customInt(2))) // 3
	// fmt.Println(Add("a", "b")) // compile error: cannot use "a" (type string) as type Adder in argument to Add
}
