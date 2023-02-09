package main

import "fmt"

type Node struct {
	Next  *Node
	Value any
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func main() {
	l := &LinkedList{}
	l.append(1)
	l.append(2)
	l.append(3)
	l.prepend(4)
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
	el := &Node{
		Value: value,
	}
	if l.Length == 0 {
		l.Head = el
	} else {
		l.Tail.Next = el
	}
	l.Tail = el
	l.Length++
}

func (l *LinkedList) prepend(value any) {
	el := &Node{
		Value: value,
	}
	if l.Length == 0 {
		l.Tail = el
	} else {
		el.Next = l.Head
	}
	l.Head = el
	l.Length++
}

func (l *LinkedList) insert(index int, value any) {
	if index >= l.Length {
		l.append(value)
	}
	if index < l.Length {
		l.prepend(value)
	}
	el := &Node{
		Value: value,
	}
	curr := l.Head
	for i := 0; i < l.Length; i++ {
		if i == index-1 {
			el.Next = curr.Next
			curr.Next = el
		}
		curr = curr.Next
	}
}

func (l *LinkedList) remove(index int) {
	if index >= l.Length {
		fmt.Println("Index out of range")
	}

	// if there is only 1 node then set head and tail equal to nil
	if l.Length == 1 {
		l.Head = nil
		l.Tail = nil
		l.Length = 0
		return
	}
	// if the index 0, then set the head to the next node
	if index == 0 {
		head := l.Head
		l.Head = head.Next
		head.Next = nil
		l.Length -= 1
		return
	}

	curr := l.Head

	for i := 0; i < l.Length; i++ {
		if i == index-1 {
			nextNode := curr.Next
			curr.Next = nextNode.Next
			nextNode.Next = nil
			if index == l.Length-1 {
				l.Tail = curr
			}
			break
		}
		curr = curr.Next

	}
	l.Length -= 1

}
