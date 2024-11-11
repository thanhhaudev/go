package main

import "fmt"

type Item[T any] interface {
	Get() T
	Set(T) Item[T]
}

type StringItem struct {
	value string
}

func (s *StringItem) Get() string {
	return s.value
}

func (s *StringItem) Set(value string) Item[string] {
	s.value = value
	return s
}

type IntItem struct {
	value int
}

func (i *IntItem) Get() int {
	return i.value
}

func (i *IntItem) Set(value int) Item[int] {
	i.value = value
	return i
}

func PrintItem[T any](item Item[T]) {
	fmt.Println(item.Get())
}

func main() {
	stringItem := &StringItem{}
	stringItem.Set("hello")
	PrintItem(stringItem)

	intItem := &IntItem{}
	intItem.Set(10)
	PrintItem(intItem)
}

/*
Advantages of Returning an Interface:
Abstraction: Hides the implementation details and exposes only the necessary methods.
Flexibility: Allows different implementations to be returned, making the code more extensible.
Decoupling: Reduces dependencies between components, making the code easier to maintain and test.

Disadvantages of Returning an Interface:
Performance: May introduce a slight overhead due to dynamic dispatch.
Type Safety: Less type information at compile time, which can lead to runtime errors if not handled carefully.
*/
