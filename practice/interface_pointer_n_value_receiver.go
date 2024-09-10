package main

/* Notes:
- In Go, Interfaces always pass by value, not by reference.
Because interfaces are wrappers around concrete types, they are reference types.
No need to use pointers to interfaces to change the underlying value.

- The value receiver is used when you want to copy the value of the receiver.
- The pointer receiver is used when you want to copy the reference of the receiver or when you want to change data in the receiver.
*/

import "fmt"

// In this Code:
// Dog implements the Animal interface with a value receiver, so both Dog and *Dog can be assigned to animal.
// Cat implements the Animal interface with a pointer receiver, so only *Cat can be assigned to animal

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string { // value receiver
	return "Woof!"
}

type Cat struct {
	Name string
}

func (c *Cat) Speak() string { // pointer receiver
	return "Meow!"
}

func main() {
	var animal Animal

	// Why Dog and *Dog are both assignable to the Animal interface?
	// The Dog type implements the Animal interface with a value receiver, so both Dog and *Dog can be assigned to the animal variable.
	animal = Dog{"Rover"}
	fmt.Println(animal.Speak())

	animal = &Dog{"Buddy"}
	fmt.Println(animal.Speak())

	// Only *Cat can be assigned to animal
	animal = &Cat{"Misty"}
	fmt.Println(animal.Speak())

	// The error message indicates that you are trying to assign a value of type Cat to a variable of type Animal,
	// but Cat does not implement the Animal interface because the Speak method has a pointer receiver.

	// In Go, for a type to implement an interface, it must implement all the methods of that interface.
	// Since the Speak method for Cat has a pointer receiver (*Cat), only *Cat (pointer to Cat) implements the Animal interface, not Cat itself.

	// animal = Cat{"Tom"} // cannot use Cat{…} (value of type Cat) as Animal value in assignment: Cat does not implement Animal (method Speak has pointer receiver)
	// fmt.Println(animal.Speak())
}
