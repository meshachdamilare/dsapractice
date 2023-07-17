package main

import "fmt"

type Node struct {
	Next  *Node
	Value int
}

type Stack struct {
	Top    *Node
	Length int
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}
func main() {
	s := Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	s.Print()
	fmt.Println("length: ", s.Length)
	s.Pop()
	s.Pop()
	s.Print()
	fmt.Println("length: ", s.Length)

}

func (s *Stack) Push(value int) {
	n := NewNode(value)
	if s.Length == 0 {
		s.Top = n
		s.Length++
		return
	}
	n.Next = s.Top
	s.Top = n
	s.Length++
}

func (s *Stack) Pop() {
	if s.Length == 0 {
		return
	}
	if s.Length == 1 {
		s.Top = nil
		s.Length--
	}
	temp := s.Top
	s.Top = s.Top.Next
	temp.Next = nil
	s.Length--
}

// Print the value in a linked singly_linked-list
func (s *Stack) Print() {
	curr := s.Top
	var list []int
	for curr != nil {
		list = append(list, curr.Value)
		curr = curr.Next
	}
	fmt.Println(list)
}
