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
// func (s *Stack[T]) isEmpty() bool {
// 	if len(s.StackList) == 0 {
// 		return true
// 	}
// 	return false
// }

func (s *Stack[T]) push(val T) {
	s.StackList = append(s.StackList, val)
}

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

func (s *Stack[T]) reverseString(word string) string {
	ns := new(Stack[rune])
	for _, v := range word {
		ns.push(v)
	}
	var reversedName string
	for len(ns.StackList) > 0 {
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

func main() {
	var s Stack[string]
	s.push("h")
	s.push("e")
	s.push("l")
	s.push("l")
	s.push("o")
	s.print()
	fmt.Println("Stack length: ", s.length())

	// Pop stack
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	s.print()
	fmt.Println("Stack length: ", s.length())

	// Reverse string through stack
	res := s.reverseString("Psalm")
	fmt.Println(res)

	res2 := s.isBalanceddParenthesis("(()())")
	fmt.Println(res2)

}
