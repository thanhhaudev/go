package main

import "fmt"

func printMap[K comparable, V any](m map[K]V) {
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func main() {
	m := map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 8,
	}

	printMap(m)

	n := map[int]string{
		1: "apple",
		2: "banana",
		3: "cherry",
	}

	printMap(n)
}
