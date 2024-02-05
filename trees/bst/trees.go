package main

import (
	"container/list"
	"fmt"
)

type Node[T int | float32] struct {
	Value       T
	Left, Right *Node[T]
}

type BST[T int | float32] struct {
	Root *Node[T]
}

func NewNode[T int | float32](value T) *Node[T] {
	return &Node[T]{
		Value: value,
	}
}

func (bt *BST[T]) insert(value T) {
	nn := NewNode(value)
	if bt.Root == nil {
		bt.Root = nn
		return
	}
	temp := bt.Root

	for {
		if value == temp.Value {
			return
		} else if value < temp.Value {
			if temp.Left == nil {
				temp.Left = nn
				return
			}
			temp = temp.Left
		} else {
			if temp.Right == nil {
				temp.Right = nn
				return
			}
			temp = temp.Right
		}
	}
}

func (bt *BST[T]) contains(value T) bool {
	if bt.Root == nil {
		return false
	}
	temp := bt.Root
	for temp != nil {
		if value < temp.Value {
			temp = temp.Left
		} else if value > temp.Value {
			temp = temp.Right
		} else {
			return true
		}
	}
	return false
}

func (bt *BST[T]) minValueNode() T {
	var res T
	if bt.Root == nil {
		return res
	}
	temp := bt.Root

	for temp != nil {
		if temp.Left == nil {
			return temp.Value
		}
		temp = temp.Left
	}
	return temp.Value
}

func (bt *BST[T]) maxValueNode() T {
	var res T
	if bt.Root == nil {
		return res
	}
	temp := bt.Root

	for temp != nil {
		if temp.Right == nil {
			return temp.Value
		}
		temp = temp.Right
	}
	return temp.Value
}

func BreadthFirstSearch[T int | float32](root *Node[T]) []T {
	var res []T
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		node := queue.Front().Value.(*Node[T])
		res = append(res, node.Value)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
		queue.Remove(queue.Front())
	}
	return res
}

func InOrder[T int | float32](node *Node[T], res []T) []T {
	if node != nil {
		res = InOrder[T](node.Left, res)
		res = append(res, node.Value)
		res = InOrder[T](node.Right, res)

	}
	return res
}

func PreOrder[T int | float32](node *Node[T], res []T) []T {
	if node != nil {
		res = append(res, node.Value)
		res = PreOrder[T](node.Left, res)
		res = PreOrder[T](node.Right, res)
	}
	return res
}

func PostOrder[T int | float32](node *Node[T], res []T) []T {
	if node != nil {
		res = PostOrder[T](node.Left, res)
		res = PostOrder[T](node.Right, res)
		res = append(res, node.Value)
	}
	return res
}

func (bt *BST[T]) printBSF() {
	result := BreadthFirstSearch[T](bt.Root)
	fmt.Println(result)
}
func (bt *BST[T]) printInOrder() {
	result := InOrder[T](bt.Root, []T{})
	fmt.Println(result)
}
func (bt *BST[T]) printPreOrder() {
	result := PreOrder[T](bt.Root, []T{})
	fmt.Println(result)
}
func (bt *BST[T]) printPostOrder() {
	result := PostOrder[T](bt.Root, []T{})
	fmt.Println(result)
}

func main() {

	// initialize two bst
	bt1 := &BST[int]{}
	bt2 := &BST[int]{}

	// insert the values below into the two trees respectively
	slice1 := []int{1, 2, 3, 4, 5, 6, 7}
	slice2 := []int{47, 21, 76, 18, 27, 52, 82}

	for _, v := range slice1 {
		bt1.insert(v)
	}

	for _, v := range slice2 {
		bt2.insert(v)
	}

	bt2.printBSF()
	bt2.printInOrder()
	bt2.printPreOrder()
	bt2.printPostOrder()

	//CHECK if a tree contains a given value
	// out := bst.contains(30)
	// fmt.Println(out)

	//FIND Max value Node of a given tree
	// out2 := bst.maxValueNode()
	// fmt.Println(out2)

}
