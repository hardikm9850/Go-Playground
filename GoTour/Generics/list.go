//I n addition to generic functions, Go also supports generic types. 
// A type can be parameterized with a type parameter, which could be useful for implementing generic data structures.
package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	value  T
}

func createList[T any](initialVal T) List[T] {
	newNode := List[T]{nil, initialVal}
	return newNode
}

func (list *List[T]) addNode(value T) {
	node := list
	for node.next != nil {
		node = node.next
	}
	newNode := List[T] {nil, value}
	node.next = &newNode
}

func (list *List[T]) print() {
	for temp := &list; ; {
		fmt.Printf("%v -> ", (*temp).value)
		if (*temp).next != nil {
			*temp = (*temp).next
		} else {
			break
		}
	}
}

func main() {
	myList := createList[int](5)
	
	myList.addNode(10)
	myList.addNode(15)
	myList.addNode(20)
	
	myList.print()
	
}
