package main

import (
	"fmt"
	"math"
)

type Node struct {
	value      int16 `json:"value"`
	leftChild  *Node `json:"leftChild"`
	rightChild *Node `json:"rightChild`
}

type Tree struct {
	root *Node
}

func makeNode(value int16) *Node {
	return &Node{value: value}
}

func (tree *Tree) insert(value int16) {
	node := makeNode(value)

	if tree.root == nil {
		tree.root = node
		return
	}

	current := tree.root
	for {
		if value < current.value {
			if current.leftChild == nil {
				current.leftChild = node
				break
			}
			current = current.leftChild
		} else {
			if current.rightChild == nil {
				current.rightChild = node
				break
			}
			current = current.rightChild
		}
	}
}

func (tree *Tree) find(value int16) bool {
	current := tree.root
	for current != nil {
		if value < current.value {
			current = current.leftChild
		} else if value > current.value {
			current = current.rightChild
		} else {
			return true
		}
	}
	return false
}

func (tree *Tree) traversePreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.value)
	tree.traversePreOrder(node.leftChild)
	tree.traversePreOrder(node.rightChild)
}

func (tree *Tree) traverseInOrder(node *Node) {
	if node == nil {
		return
	}
	tree.traversePreOrder(node.leftChild)
	fmt.Println(node.value)
	tree.traversePreOrder(node.rightChild)
}

func (tree *Tree) traversePostOrder(node *Node) {
	if node == nil {
		return
	}
	tree.traversePreOrder(node.leftChild)
	tree.traversePreOrder(node.rightChild)
	fmt.Println(node.value)
}

func (tree *Tree) height(root *Node) float64 {
	if root == nil {
		return -1
	}
	if root.leftChild == nil && root.rightChild == nil {
		return 0
	}
	return 1 + math.Max(tree.height(root.leftChild), tree.height(root.rightChild))
}

func (tree *Tree) min(root *Node) float64 {
	left := min(root.leftChild)
	right := min(root.rightChild)

	return math.Min(math.Min(left, right), float64(root.value))
}

func factorial(value int16) int16 {
	if value == 0 {
		return 1
	}
	return value * factorial(value-1)
}

func main() {
	fmt.Println("Binary tree")
	tree := Tree{}
	fmt.Println("Insert in tree")
	tree.insert(7)
	tree.insert(4)
	tree.insert(9)
	tree.insert(1)
	tree.insert(6)
	tree.insert(8)
	tree.insert(10)
	fmt.Println("Find in tree")
	result := tree.find(13)
	fmt.Printf("%t\n", result)
	fmt.Println("Pre order traversal")
	tree.traversePreOrder(tree.root)
	fmt.Println("In order traversal")
	tree.traverseInOrder(tree.root)
	fmt.Println("Post order traversal")
	tree.traversePostOrder(tree.root)
	fmt.Println("Tree Height")
	fmt.Printf("%g", tree.height(tree.root))
}
