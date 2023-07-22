package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Next  *Node
	Prev  *Node
	Value int
}

type DoublyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func main() {
	l := DoublyLinkedList{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Push(5)
	l.Push(6)
	l.Print()
	fmt.Println("list length:", l.Length)

	// ReverseBetween
	/*l.ReverseBetween(0, 5)
	l.Print()*/

	// FindKthFromEnd
	/*	n := l.FindKthFromEnd(4)
		PrintNodeValue(n)*/

	// findMiddleNode example
	/*n := l.FindMiddleNode()
	PrintNodeValue(n)*/

	// Pop, Shift and Unshift example
	/*	fmt.Println("length: ", l.Length)
		l.Pop()
		l.Print()
		fmt.Println("length: ", l.Length)
		l.Unshift(12)
		l.Print()
		fmt.Println("length: ", l.Length)
		l.Shift()
		l.Print()
		fmt.Println("length: ", l.Length)*/

	// GetNodeValueByIndex example
	/*	n, err := l.GetNodeValueByIndex(2)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			PrintNodeValue(n)
		}*/

	// SetNodeValueByIndex example
	/*	err := l.SetNodeValueByIndex(33, 2)
		if err != nil {
			fmt.Println(err.Error())
		}
		l.Print()*/

	// insert example
	/*	err := l.Insert(22, 2)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			l.Print()
			fmt.Println("linked-list length:", l.Length)
		}*/

	// remove example
	err := l.Remove(3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		l.Print()
		fmt.Println("linked-list length:", l.Length)
	}

	// reverse example
	/*l.Reverse()
	l.Print()*/
}

func (l *DoublyLinkedList) Push(value int) {
	n := NewNode(value)
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		l.Length++
		return
	}
	l.Tail.Next = n
	n.Prev = l.Tail
	l.Tail = n
	l.Length++
}

func (l *DoublyLinkedList) Pop() {
	if l.Length == 0 {
		fmt.Println("cant pop from empty linked-list")
		return
	}
	if l.Length == 1 {
		val := l.Head.Value
		l.Head = nil
		l.Tail = nil
		l.Length--
		fmt.Println(val)
		return
	}
	temp := l.Tail
	l.Tail = l.Tail.Prev
	l.Tail.Next = nil
	l.Length--
	fmt.Println("Pop: ", temp.Value)
	return
}

func (l *DoublyLinkedList) Unshift(value int) {
	n := NewNode(value)
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		l.Length++
		return
	}
	n.Next = l.Head
	l.Head.Prev = n
	l.Head = n
	l.Length++
}

func (l *DoublyLinkedList) Shift() {
	if l.Head == nil {
		fmt.Println("Cannot remove from empty list")
		return
	}
	if l.Length == 1 {
		l.Head = nil
		l.Tail = nil
		l.Length--
	}
	temp := l.Head
	l.Head = temp.Next
	l.Head.Prev = nil
	l.Length--
}
func (l *DoublyLinkedList) GetNodeValueByIndex(idx int) (*Node, error) {
	if idx < 0 || idx > l.Length {
		return nil, errors.New("index out of range")
	}
	curr := l.Head
	for i := 0; i < idx-1; i++ {
		curr = curr.Next
	}
	return curr, nil
}

func (l *DoublyLinkedList) SetNodeValueByIndex(value, idx int) error {
	n, err := l.GetNodeValueByIndex(idx)
	if err != nil {
		return err
	}
	n.Value = value
	return nil
}

func (l *DoublyLinkedList) Insert(value, idx int) error {
	n := NewNode(value)
	if idx == 0 {
		l.Unshift(value)
		return nil
	}
	if idx == l.Length {
		l.Push(value)
		return nil
	}
	if idx < 0 || idx > l.Length {
		return errors.New("index out of range")
	}
	before, err := l.GetNodeValueByIndex(idx - 1)
	if err != nil {
		return err
	}
	n.Next = before.Next
	before.Next.Prev = n
	before.Next = n
	n.Prev = before
	l.Length++
	return nil
}

func (l *DoublyLinkedList) Remove(idx int) error {
	if idx < 0 || idx >= l.Length {
		return errors.New("index out of range")
	}
	if idx == 0 {
		l.Shift()
		return nil
	}
	if idx == l.Length-1 {
		l.Pop()
		return nil
	}
	before, err := l.GetNodeValueByIndex(idx - 1)
	if err != nil {
		return err
	}
	temp := before.Next
	before.Next = temp.Next
	temp.Next.Prev = before
	temp.Prev = nil
	temp.Next = nil
	return nil
}

// Print the value in a linked singly_linked-list
func (l *DoublyLinkedList) Print() {
	curr := l.Head
	list := make([]int, 0)
	for curr != nil {
		list = append(list, curr.Value)
		curr = curr.Next
	}
	fmt.Println(list)
}

// PrintNodeValue print the value at the current node
func PrintNodeValue(n *Node) {
	fmt.Println(n.Value)
}
