package main

import (
	"fmt"
	"sync"
)

type BinaryTree struct {
	sync.RWMutex
	root *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func main() {

	var (
		inOrderTraversal   []int
		preOrderTraversal  []int
		postOrderTraversal []int
	)

	input := []int{2, 4, 19, 7, 30}
	binaryTree := new(BinaryTree)

	for _, value := range input {
		binaryTree.InsertNode(value)
	}

	binaryTree.Print()

	inOrder(binaryTree.root, &inOrderTraversal)
	fmt.Println("\nInOrder traversal-----> ", inOrderTraversal)

	preOrder(binaryTree.root, &preOrderTraversal)
	fmt.Println("\nPreOrder traversal-----> ", preOrderTraversal)

	postOrder(binaryTree.root, &postOrderTraversal)
	fmt.Println("\nPostOrder traversal-----> ", postOrderTraversal)
}

func newNode(value int) *Node {
	return &Node{
		value: value,
	}
}

//InsertNode will add value into the binary tree
func (t *BinaryTree) InsertNode(value int) {
	t.RWMutex.Lock()
	defer t.RWMutex.Unlock()

	if t.root == nil {
		t.root = newNode(value)
		return
	}
	insert(value, t.root)
}

func insert(value int, node *Node) {

	if value < node.value {

		if node.left == nil {
			node.left = newNode(value)
		} else {
			insert(value, node.left)
		}
	} else if value > node.value {

		if node.right == nil {
			node.right = newNode(value)
		} else {
			insert(value, node.right)
		}
	}
}

func inOrder(node *Node, array *[]int) {

	if node.left != nil {
		inOrder(node.left, array)
	}

	*array = append(*array, node.value)

	if node.right != nil {
		inOrder(node.right, array)
	}
}

func preOrder(node *Node, array *[]int) {

	*array = append(*array, node.value)

	if node.left != nil {
		preOrder(node.left, array)
	}
	if node.right != nil {
		preOrder(node.right, array)
	}
}

func postOrder(node *Node, array *[]int) {

	if node.left != nil {
		postOrder(node.left, array)
	}
	if node.right != nil {
		postOrder(node.right, array)
	}
	*array = append(*array, node.value)
}

func (t *BinaryTree) Print() {
	t.RWMutex.RLock()
	defer t.RWMutex.RUnlock()

	fmt.Print("----------BINARY TREE --------------\n", " M-> ROOT-NODE; L-> LEVEL; Ln-> LEFT-NODE; Rn-> RIGHT-NODE \n")
	print(t.root, 0, 0, "R")
	fmt.Println("----------------------------------------------------")
}

func print(node *Node, space, level int, nodeType string) {
	if node == nil {
		return
	}

	fmt.Printf(" L:%d", level)
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("   - %s: %v\n", nodeType, node.value)
	level++
	print(node.left, space+4, level, "Ln")
	print(node.right, space+4, level, "Rn")
}
