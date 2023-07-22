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
	l.Push(1)
	l.Push(4)
	l.Push(3)
	l.Push(2)
	l.Push(5)
	l.Push(2)
	l.Push(4)
	l.Print()

	// Remove duplicates
	l.RemoveDuplicates()
	l.Print()
	// Partition List
	/*	l.PartitionList(3)
		l.Print()*/

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
		fmt.Println("linked-list length:", l.Length)*/

	// remove example
	/*	err := l.Remove(3)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			l.Print()
			fmt.Println("linked-list length:", l.Length)
		}*/

	// reverse example
	/*	l.Reverse()
		l.Print()
	*/

}

// Push Append to the end of a linked-linked-list
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

// Pop remove the last element of a linked-list
func (l *LinkedList) Pop() {
	if l.Length == 0 {
		fmt.Println("cant pop from empty linked-list")
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

// Unshift Adds element to beginning of the linked-list
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

// Shift Remove the first node of a linked-list
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

// Reverse reverses a linked list
func (l *LinkedList) Reverse() {
	curr := l.Head
	var prev *Node
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	l.Head = prev
	return
}

/*Implement a member function called findMiddleNode() that finds and returns the middle node of the linked-list.
Note: this LinkedList implementation does not have a length member variable.

Output:
Return the middle node of the linked-list.
If the linked-list has an even number of nodes, return the second middle node (the one closer to the end).

Constraints:
You are not allowed to use any additional data structures (such as arrays) or modify the existing data structure.
You can only traverse the linked-list once.*/

// FindMiddleNode to get the middle node assuming the length of the linkedlist is not known
func (l *LinkedList) FindMiddleNode() *Node {
	slow := l.Head
	fast := l.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

/*LL: Find Kth Node From End ( ** Interview Question)
Implement a member function called findKthFromEnd(k) that finds and returns the kth node from the end of the linked list.
Note: this LinkedList implementation does not have a length member variable.

Output:
Return the kth node from the end of the linked list.

If the value of k is greater than or equal to the number of nodes in the list, return null.
Constraints:
You are not allowed to use any additional data structures (such as arrays) or modify the existing data structure.
You can only traverse the linked list once.

Example 1:
Suppose you have a LinkedList object, list, with the following values:
1 -> 2 -> 3 -> 4 -> 5
After calling the findKthFromEnd(2) function:
let kthNode = list.findKthFromEnd(2);
The kthNode should have the value 4.

Example 2:
Now suppose you have a LinkedList object, list, with the following values: 1 -> 2 -> 3 -> 4 -> 5 -> 6
After calling the findKthFromEnd(4) function:
let kthNode = list.findKthFromEnd(4);
The kthNode should have the value 3.
*/

func (l *LinkedList) FindKthFromEnd(k int) *Node {
	fast := l.Head
	slow := l.Head

	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

/*
LL: Reverse Between ( ** Interview Question)
Implement a member function called reverseBetween(m, n) that reverses the nodes between indexes (using 0-based indexing)  m and n (inclusive) in the linked list.
Note: this linked list class does not have a tail which will make this method easier to implement.

Output:
The function should reverse the nodes between the indexes m and n in the linked list. The reversal should be done in-place.

Constraints:
You are not allowed to use any additional data structures (such as arrays) or modify the existing data structure.
You can only traverse the linked list once.

Example 1:
Suppose you have a LinkedList object, list, with the following values:
1 -> 2 -> 3 -> 4 -> 5

After calling the reverseBetween(1, 3) function:
list.reverseBetween(1, 3);
The linked list should now have the following values:
1 -> 4 -> 3 -> 2 -> 5

Example 2:
Now suppose you have a LinkedList object, list, with the following values:
1 -> 2 -> 3 -> 4 -> 5 -> 6

After calling the reverseBetween(3, 5) function:
list.reverseBetween(3, 5);
The linked list should now have the following values:
1 -> 2 -> 3 -> 6 -> 5 -> 4
*/
func (l *LinkedList) ReverseBetween(m, n int) {
	if l.Head == nil {
		return
	}
	dummy := new(Node)
	dummy.Next = l.Head
	prev := dummy

	for i := 0; i < m; i++ {
		prev = prev.Next
	}
	curr := prev.Next

	for i := 0; i < n-m; i++ {
		temp := curr.Next
		curr.Next = temp.Next
		temp.Next = prev.Next
		prev.Next = temp
	}
	l.Head = dummy.Next
}

func (l *LinkedList) PartitionList(x int) {
	if l.Head == nil || l.Head.Next == nil {
		return
	}
	beforeHead := new(Node)
	afterHead := new(Node)
	before := beforeHead
	after := afterHead
	curr := l.Head
	for curr != nil {
		if curr.value < x {
			before.Next = curr
			before = curr
		} else {
			after.Next = curr
			after = curr
		}
		curr = curr.Next
	}

	after.Next = nil
	before.Next = afterHead.Next

	l.Head = beforeHead.Next

	return
}

func (l *LinkedList) RemoveDuplicates() {
	uniqueMap := make(map[int]bool)
	prev := new(Node)
	curr := l.Head
	for curr != nil {
		if _, found := uniqueMap[curr.value]; found {
			prev.Next = curr.Next
		} else {
			uniqueMap[curr.value] = true
			prev = curr
		}
		curr = curr.Next
	}
}

// Print the value in a linked singly_linked-list
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
