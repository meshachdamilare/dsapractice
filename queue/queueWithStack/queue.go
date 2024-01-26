package main

import "fmt"

type Stack[T any] struct {
	StackList []T
}

func (s *Stack[T]) push(val T) {
	s.StackList = append(s.StackList, val)
}

func (s *Stack[T]) pop() (T, bool) {
	var res T
	l := len(s.StackList)
	if l == 0 {
		return res, false
	}
	res = s.StackList[l-1]
	s.StackList = s.StackList[:l-1]
	return res, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.StackList) == 0
}

type MyQueue[T any] struct {
	queueList *Stack[T]
	opStack   *Stack[T]
}

func (q *MyQueue[T]) Enqueue(val T) {
	for !q.queueList.isEmpty() {
		if top, ok := q.queueList.pop(); ok {
			q.opStack.push(top)
		}
	}
	if q.queueList.isEmpty() {
		q.queueList.push(val)
	}
	for !q.opStack.isEmpty() {
		if top, ok := q.opStack.pop(); ok {
			q.queueList.push(top)
		}
	}
}
func (q *MyQueue[T]) Dequeue() (T, bool) {

	for !q.queueList.isEmpty() {
		if len(q.queueList.StackList) == 1 {
			break
		}
		if top, ok := q.queueList.pop(); ok {
			q.opStack.push(top)
		}
	}

	res, ok := q.queueList.pop()

	if q.queueList.isEmpty() {
		for !q.opStack.isEmpty() {
			if top, ok := q.opStack.pop(); ok {
				q.queueList.push(top)
			}
		}
	}
	return res, ok

}

func (q *MyQueue[T]) print() {
	fmt.Println(q.queueList)
}

func main() {
	q := &MyQueue[int]{
		queueList: &Stack[int]{},
		opStack:   &Stack[int]{},
	}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	q.print()

	res1, _ := q.Dequeue()
	fmt.Println(res1)
	q.print()

	res2, _ := q.Dequeue()
	fmt.Println(res2)
	q.print()
}
