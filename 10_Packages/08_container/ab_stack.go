package main

import (
	"container/list"
	"fmt"
	"reflect"
)

// Stack represents a LIFO (Last-In-First-Out) data structure.
type Stack struct {
	list *list.List
}

func (s *Stack) Push(value interface{}) {
	s.list.PushFront(value)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil // Return nil if the stack is empty
	}
	element := s.list.Front()
	s.list.Remove(element)
	return element.Value
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil // Return nil if the stack is empty
	}
	return s.list.Front().Value
}

func (s *Stack) IsEmpty() bool {
	return s.list.Len() == 0
}

func (s *Stack) Size() int {
	return s.list.Len()
}

func (s *Stack) Print() {
	for element := s.list.Front(); element != nil; element = element.Next() {
		fmt.Print(element.Value, " ")
	}
	fmt.Println()
}

func NewStack() *Stack {
	return &Stack{list: list.New()}
}

func main() {
	// Create a new stack
	stack := NewStack()
	fmt.Printf("Type of stack: %T\n", stack)

	// Push elements onto the stack
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")

	// Print the stack
	fmt.Print("Stack after pushes: ")
	stack.Print()

	// Peek at the top element
	fmt.Println("Peek:", stack.Peek())

	// Pop elements from the stack
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Pop:", stack.Pop())

	// Print the stack
	fmt.Print("Stack after pops: ")
	stack.Print()

	// Check if the stack is empty
	fmt.Println("Is stack empty?", stack.IsEmpty())

	// Get the size of the stack
	fmt.Println("Stack size:", stack.Size())

	// Use reflection to access the unexported field "len" in the list
	fv := reflect.ValueOf(stack.list).Elem().FieldByName("len")
	fmt.Println("Reflection - List length:", fv.Int())

}
