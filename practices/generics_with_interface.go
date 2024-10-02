package main

import "fmt"

type Item[T any] interface {
	Get() T
	Set(T) T
}

type StringItem struct {
	value string
}

func (s *StringItem) Get() string {
	return s.value
}

func (s *StringItem) Set(value string) *StringItem {
	s.value = value
	return s
}

type IntItem struct {
	value int
}

func (i *IntItem) Get() int {
	return i.value
}

func (i *IntItem) Set(value int) *IntItem {
	i.value = value
	return i
}

func main() {
	stringItem := &StringItem{}
	stringItem.Set("hello")
	fmt.Println(stringItem.Get())

	intItem := &IntItem{}
	intItem.Set(10)
	fmt.Println(intItem.Get())
}
