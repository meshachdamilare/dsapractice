package main

import "fmt"

type Node struct {
	Next  *Node
	Prev  *Node
	Value any
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func main() {
	l := &LinkedList{}

	l.prepend(4)
	l.prepend(3)
	l.prepend(2)
	l.prepend(1)
	l.remove(1)
	l.display()
}

func (l *LinkedList) display() {
	curr := l.Head
	for curr != nil {
		fmt.Println(curr.Value)
		curr = curr.Next
	}
}

func (l *LinkedList) append(value any) {
	n := &Node{
		Value: value,
	}
	if l.Length == 0 {
		l.Head = n
	} else {
		l.Tail.Next = n
		n.Prev = l.Tail
	}
	l.Tail = n
	l.Length++
}

func (l *LinkedList) prepend(value any) {

}

func (l *LinkedList) insert(index int, value any) {
	if index == 0 {
		l.prepend(value)
		return
	}
	if index >= l.Length {
		l.append(value)
		return
	}
	n := &Node{
		Value: value,
	}
	median := l.Length / 2
	var header *Node
	if median > index {
		header = l.transverseToIndexFromHead(index - 1)
	} else {
		header = l.transverseToIndexFromTail(index - 1)
	}
	n.Prev = header
	n.Next = header.Next
	header.Next.Prev = n
	header.Next = n
	l.Length++
}

func (l *LinkedList) remove(index int) {
	if index >= l.Length {
		fmt.Println("Index out of bounds")
		return
	}
	if index == 0 {
		head := l.Head
		head.Next.Prev = nil
		l.Head = l.Head.Next
		head.Next = nil
		l.Length--
		return
	}
	if index == l.Length-1 {
		tail := l.Tail
		tail.Prev.Next = nil
		l.Tail = l.Tail.Prev
		tail.Prev = nil
		l.Length--
		return
	}
	median := l.Length / 2
	var curr *Node
	if median > index {
		curr = l.transverseToIndexFromHead(index)
	} else {
		curr = l.transverseToIndexFromTail(index)
	}
	curr.Prev.Next = curr.Next
	curr.Next.Prev = curr.Prev
	curr.Next = nil
	curr.Prev = nil
	l.Length--
}

func (l *LinkedList) transverseToIndexFromHead(index int) *Node {
	counter := 0
	curr := l.Head
	for counter != index {
		curr = curr.Next
		counter++
	}
	return curr
}

func (l *LinkedList) transverseToIndexFromTail(index int) *Node {
	counter := l.Length - 1
	curr := l.Tail
	for counter != index {
		curr = curr.Prev
		counter--
	}
	return curr

}
