package main

import "fmt"

// Maps are a collection of key-value pairs. The keys are unique within a map, while the values may not be unique.
// The syntax to create a map is map[keyType]valueType. The keyType can be any type that supports the equality operator, ==. (comparable types)
// The zero value of a map is nil. If you try to add elements to a nil map, a runtime panic will occur.
// A nil map is equivalent to an empty map, but you cannot add elements to a nil map, the length of a nil map is 0, and the capacity of a nil map is 0.

func main() {
	var nilMap map[string]int
	// nilMap is declared to be a map with string keys and int values, but it is nil.
	fmt.Println(nilMap == nil)       // true
	fmt.Printf("nilMap: %v", nilMap) // nilMap: map[]
	// nilMap["one"] = 1 // panic: assignment to entry in nil map

	// We can use a := declaration to create a map variable by assigning a map literal to it.
	nonNilMap := map[string]int{}            // using an empty map literal, we also can provide initial key-value pairs. e.g. map[string]int{"one": 1, "two": 2}
	fmt.Println(nonNilMap == nil)            // false
	fmt.Printf("nonNilMap: %v\n", nonNilMap) // nonNilMap: map[]

	nonNilWithInitialValues := map[string][]string{
		"fruits":  []string{"apple", "banana", "orange"},
		"numbers": []string{"one", "two", "three"},
	}

	fmt.Println(nonNilWithInitialValues == nil)                          // false
	fmt.Printf("nonNilWithInitialValues: %v\n", nonNilWithInitialValues) // nonNilWithInitialValues: map[fruits:[apple banana orange] numbers:[one two three]
	// the key followed by a colon and then the value, and each key-value pair is separated by a comma.
	// the command separating each key-value pair in the map, even on the last pair.

	// => different between nil and empty map is that we can add elements to an empty map, but we cannot add elements to a nil map.
}
