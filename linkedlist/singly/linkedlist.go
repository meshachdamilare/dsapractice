package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	Next  *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
	}
}

func main() {
	l := LinkedList{}
	l.Push(4)
	l.Push(5)
	l.Push(6)
	l.Push(7)
	l.Print()
	fmt.Println("list length:", l.Length)

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
	/*	n, err := l.GetNodeValueByIndex(3)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			PrintNodeValue(n)
		}*/

	// SetNodeValueByIndex example
	/*	err := l.SetNodeValueByIndex(33, 2)
		if err != nil {
			fmt.Println(err.Error())
		}*/

	// insert example
	/*	err := l.Insert(22, 2)
		if err != nil {
			return
		}
		l.Print()
		fmt.Println("list length:", l.Length)*/

	// remove example
	/*	err := l.Remove(3)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			l.Print()
			fmt.Println("list length:", l.Length)
		}*/

	// reverse example
	/*l.Reverse()
	l.Print()*/
}

// Push Append to the end of a linked-list
func (l *LinkedList) Push(value int) {
	node := NewNode(value)
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	}
	l.Tail.Next = node
	l.Tail = node
	l.Length++
}

// Pop remove the last element of a list
func (l *LinkedList) Pop() {
	if l.Length == 0 {
		fmt.Println("cant pop from empty list")
		return
	}
	if l.Length == 1 {
		val := l.Head.value
		l.Head = nil
		l.Tail = nil
		l.Length = 0
		fmt.Println(val)
	}
	curr := l.Head
	pre := l.Head
	for curr.Next != nil {
		temp := curr.Next
		pre = curr
		curr = temp
	}
	pre.Next = nil
	l.Tail = pre
	l.Length--
	return
}

// Unshift Adds element to beginning of the list
func (l *LinkedList) Unshift(value int) {
	node := NewNode(value)
	if l.Length == 0 {
		l.Head = node
		l.Tail = node
	}
	node.Next = l.Head
	l.Head = node
	l.Length++
}

// Shift Remove the first node of a list
func (l *LinkedList) Shift() {
	if l.Length == 0 {
		return
	}
	if l.Length == 1 {
		l.Head = nil
		l.Tail = nil
	}
	temp := l.Head
	l.Head = l.Head.Next
	temp.Next = nil
	l.Length--
}

// GetNodeValueByIndex Get the node value at a given index
func (l *LinkedList) GetNodeValueByIndex(idx int) (*Node, error) {
	if idx < 0 || idx > l.Length {
		return nil, errors.New("index out of range")
	}
	temp := l.Head
	for i := 0; i < idx; i++ {
		temp = temp.Next
	}
	return temp, nil
}

// SetNodeValueByIndex change the node value at the given index
func (l *LinkedList) SetNodeValueByIndex(value int, idx int) error {
	temp, err := l.GetNodeValueByIndex(idx)
	if err != nil {
		return err
	}
	temp.value = value
	return nil
}

// Insert insert a new node value at the given index
func (l *LinkedList) Insert(value int, idx int) error {
	if idx == 0 {
		l.Unshift(value)
		return nil
	}
	if idx == l.Length {
		l.Push(value)
		return nil
	}
	if idx < 0 || idx > l.Length {
		return errors.New("insert error: Index out of range")
	}
	newNode := NewNode(value)
	// gets the node at the index before the given index
	temp, _ := l.GetNodeValueByIndex(idx - 1)
	newNode.Next = temp.Next
	temp.Next = newNode
	l.Length++
	return nil
}

// Remove removes node value at the given index
func (l *LinkedList) Remove(idx int) error {
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
	// gets the node at the index before the given index
	before, _ := l.GetNodeValueByIndex(idx - 1)
	temp := before.Next
	before.Next = temp.Next
	temp.Next = nil
	l.Length--
	return nil
}

func (l *LinkedList) Reverse() {
	/*
		Steps involve
		1. Switch Head and Tail i.e. Tail becomes the Head and Head becomes the Tail
		2. Declare two variables to track the prev node and next node of the current node
		3. Swap the nodes
	*/

	// 1. Switch Head and Tail i.e. Tail becomes the Head and Head becomes the Tail
	temp := l.Head
	l.Head = l.Tail
	l.Tail = temp
	// 2. Declare two variables to track the prev node and next node of the current node
	next := temp.Next
	var prev *Node // tactics to set  prev to nil
	for i := 0; i < l.Length; i++ {
		next = temp.Next
		temp.Next = prev
		prev = temp
		temp = next
	}
	return
}

// Print the value in a linked list
func (l *LinkedList) Print() {
	curr := l.Head
	var list []int
	for curr != nil {
		list = append(list, curr.value)
		curr = curr.Next
	}
	fmt.Println(list)
}

// PrintNodeValue print the value at the current node
func PrintNodeValue(n *Node) {
	fmt.Println(n.value)
}
