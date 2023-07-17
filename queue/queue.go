package main

import "fmt"

type Node struct {
	Next  *Node
	Value int
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

type Queue struct {
	First  *Node
	Last   *Node
	Length int
}

func main() {
	q := Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Print()
	q.Dequeue()
	q.Dequeue()
	q.Print()

}

func (q *Queue) Enqueue(value int) {
	n := NewNode(value)
	if q.Length == 0 {
		q.First = n
		q.Last = n
		q.Length++
		return
	}
	q.Last.Next = n
	q.Last = n
	q.Length++
}

func (q *Queue) Dequeue() {
	if q.Length == 0 {
		return
	}
	if q.Length == 1 {
		q.First = nil
		q.Last = nil
		q.Length--
		return
	}
	temp := q.First
	q.First = q.First.Next
	temp.Next = nil
	q.Length--
	return
}

// Print the value in a linked singly_linked-list
func (s *Queue) Print() {
	curr := s.First
	var list []int
	for curr != nil {
		list = append(list, curr.Value)
		curr = curr.Next
	}
	fmt.Println(list)
}
