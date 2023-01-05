package main

import (
	"container/list"
	"fmt"
)

type Elem struct {
	A int
	B int
}

func main() {

	e1 := Elem{A: 1, B: 2}
	e2 := Elem{A: 5, B: 4}

	mylist := list.New()
	mylist.PushBack(e1)
	mylist.PushFront(e2)

	for element := mylist.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		if (element.Value.(Elem)).A == 1 {
			mylist.Remove(element)
		}
	}
	for element := mylist.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		fmt.Println(element.Value.(Elem).A, element.Value.(Elem).B)
	}
}
