package main

import (
	"fmt"
)

type Stack[T any] struct {
	StackList []T
}

func (s *Stack[T]) getStackList() []T {
	return s.StackList
}

func (s *Stack[T]) print() {
	fmt.Println(s.StackList)
}

func (s *Stack[T]) length() int {
	return len(s.getStackList())
}

// Check if the stack is empty or nnot
func (s *Stack[T]) isEmpty() bool {
	return len(s.StackList) == 0
}

// Push into stack
func (s *Stack[T]) push(val T) {
	s.StackList = append(s.StackList, val)
}

// Remove the top of a stack, true and false if the top is empty
// It modifies th stack
func (s *Stack[T]) pop() (T, bool) {
	var res T
	l := s.length()
	if l == 0 {
		return res, false
	}
	res = s.StackList[l-1]
	s.StackList = s.StackList[:l-1]
	return res, true
}

// Check out the top of a stack, true and false if the top is empty
// It does not modify the stack
func (s *Stack[T]) peek() (T, bool) {
	var res T
	l := s.length()
	if l == 0 {
		return res, false
	}
	return s.StackList[l-1], true
}

func (s *Stack[T]) reverseString(word string) string {
	ns := new(Stack[rune])
	for _, v := range word {
		ns.push(v)
	}
	var reversedName string
	for !ns.isEmpty() {
		if ch, ok := ns.pop(); ok {
			reversedName += string(ch)
		}
	}
	return reversedName
}

func (s *Stack[T]) isBalanceddParenthesis(p string) bool {
	ns := new(Stack[rune])
	for _, v := range p {
		switch v {
		case '(', '{', '[':
			ns.push(v)
		case ')':
			if top, ok := ns.pop(); !ok || top != '(' {
				return false
			}
		case ']':
			if top, ok := ns.pop(); !ok || top != '[' {
				return false
			}
		case '}':
			if top, ok := ns.pop(); !ok || top != '{' {
				return false
			}
		}

	}
	if len(ns.StackList) == 0 {
		return true
	} else {
		return false
	}
}

func sortStack[T int](s *Stack[T]) {
	additionalStack := new(Stack[T])
	for !s.isEmpty() {
		temp, _ := s.pop()

		for !additionalStack.isEmpty() {
			top, _ := additionalStack.pop()
			if top > temp {
				s.push(top)
			} else {
				additionalStack.push(top)
				break
			}
		}
		additionalStack.push(temp)
	}
	for !additionalStack.isEmpty() {
		top, _ := additionalStack.pop()
		s.push(top)
	}
}

func main() {
	var s Stack[string]
	s.push("h")
	s.push("e")
	s.push("l")
	s.push("l")
	s.push("o")
	s.print()
	fmt.Println("Stack length: ", s.length())
	fmt.Println("Is empty: ", s.isEmpty())
	// Pop stack
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	s.print()
	fmt.Println("Stack length: ", s.length())

	// fmt.Println("Is empty: ", s.isEmpty())

	// Reverse string through stack
	res := s.reverseString("Psalm")
	fmt.Println(res)

	res2 := s.isBalanceddParenthesis("(()())")
	fmt.Println(res2)

	var s2 Stack[int]

	s2.push(2)
	s2.push(1)
	s2.push(5)
	s2.push(3)
	s2.push(4)

	s2.print()

	sortStack[int](&s2)
	s2.print()
	fmt.Println(s2.length())

}
