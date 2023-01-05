package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type StackLinkedList struct {
	head *Node
	size int
}

// Size() function
func (s *StackLinkedList) Size() int {
	return s.size
}

// IsEmpty() function
func (s *StackLinkedList) IsEmpty() bool {
	return s.size == 0
}

// Peek() function
func (s *StackLinkedList) Peek() int {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return 0
	}
	return s.head.value
}

// Push() function
func (s *StackLinkedList) Push(value int) {
	s.head = &Node{value, s.head}
	s.size++
}

// Pop() function
func (s *StackLinkedList) Pop() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return 0, false
	}
	value := s.head.value
	s.head = s.head.next
	s.size--
	return value, true
}

// Print() function
func (s *StackLinkedList) Print() {
	temp := s.head
	fmt.Print("Values stored in stack are: ")
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

func stack_test() {
	var stack StackLinkedList // create a stack variable of type StackLinkedList
	/* Adding items to stack */
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Print() // Print the stack
	fmt.Println("Checking length of stack:")
	fmt.Println(stack.Size()) // Get Length
	fmt.Println("Removing an Item:")
	stack.Pop() // Remove an item
	stack.Print()
	fmt.Println("Getting Item at Top of Linked List")
	fmt.Println(stack.Peek()) // Get Top-most item
}
