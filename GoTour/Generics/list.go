//I n addition to generic functions, Go also supports generic types. 
// A type can be parameterized with a type parameter, which could be useful for implementing generic data structures.
package main

import (
	"fmt"
)

type node[T any] struct {
	Data T
	next *node[T]
}

type list[T any] struct {
	start *node[T]
}

func (l *list[T]) add(data T) {
	n := node[T]{
		Data: data,
		next: nil,
	}

	if l.start == nil {
		l.start = &n
		return
	}

	if l.start.next == nil {
		l.start.next = &n
		return
	}

	temp := l.start
	l.start = l.start.next
	l.add(data)
	l.start = temp
}

func main() {
	var myList list[int]
	fmt.Println(myList)
	myList.add(12)
	myList.add(9)
	myList.add(3)
	myList.add(9)

	// Print all elements
	cur := myList.start
	for {
		fmt.Println("*", cur)
		if cur == nil {
			break
		}
		fmt.Println("*", cur.Data)
		cur = cur.next
	}
}

// ------- output ------- 

{<nil>}
* &{12 0xc000014080}
* 12
* &{9 0xc0000140a0}
* 9
* &{3 0xc0000140d0}
* 3
* &{9 <nil>}
* 9
* <nil>
