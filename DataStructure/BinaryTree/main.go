package main

import (
	"fmt"
	"math/rand"
	"time"

	"./myutils"
)

type node struct {
	val         int
	left, right *node
}

func (n *node) new(val int) {
	n.val = val
	n.left, n.right = nil, nil
}

type binaryTree struct {
	root *node
}

func (b *binaryTree) new() {
	b.root = nil
}
func (b *binaryTree) insert(val int) {
	var newNode node
	newNode.new(val)
	if b.root == nil {
		b.root = &newNode
		return
	}
	current := b.root
	for {
		if val == b.root.val {
			return
		}
		if val < current.val {
			if current.left == nil {
				current.left = &newNode
				return
			} else {
				current = current.left
			}
		} else {
			if current.right == nil {
				current.right = &newNode
				return
			} else {
				current = current.right
			}
		}
	}
}
func (b *binaryTree) find(val int) *node {
	if b.root == nil {
		return nil
	}
	current := b.root
	found := false
	for !found && current != nil {
		if val < current.val {
			current = current.left
		} else if val > current.val {
			current = current.right
		} else {
			found = true
		}
	}
	return current
}

// Breath Search Tree implementation
func (b *binaryTree) Bst() []int {
	var node1 *node
	node1 = b.root
	var result []int
	var queue myutils.Queue
	queue.Enqueue(node1)
	for queue.Size != 0 {
		node1 = queue.Dequeue().(*node)
		result = append(result, node1.val)
		if node1.left != nil {
			queue.Enqueue(node1.left)
		}
		if node1.right != nil {
			queue.Enqueue(node1.right)
		}
	}
	return result
}

func (b *binaryTree) DFSPreOrder() []int {
	var data []int
	current := b.root
	var traverse func(n *node)
	traverse = func(n *node) {
		data = append(data, n.val)
		if n.left != nil {
			traverse(n.left)
		}
		if n.right != nil {
			traverse(n.right)
		}
	}
	traverse(current)
	return data
}
func (b *binaryTree) DFSPostOrder() []int {
	var data []int
	current := b.root
	var traverse func(n *node)
	traverse = func(n *node) {
		if n.left != nil {
			traverse(n.left)
		}
		if n.right != nil {
			traverse(n.right)
		}
		data = append(data, n.val)
	}
	traverse(current)
	return data
}
func (b *binaryTree) DFSInOrder() []int {
	var data []int
	current := b.root
	var traverse func(n *node)
	traverse = func(n *node) {
		if n.left != nil {
			traverse(n.left)
		}
		data = append(data, n.val)
		if n.right != nil {
			traverse(n.right)
		}
	}
	traverse(current)
	return data
}
func main() {
	rand.Seed(time.Now().UnixNano())
	var btree binaryTree
	btree.new()
	for i := 0; i < 7; i++ {
		btree.insert(rand.Intn(1000))
	}
	fmt.Println(btree.Bst())
	fmt.Println(btree.DFSPreOrder())
	fmt.Println(btree.DFSPostOrder())
	fmt.Println(btree.DFSInOrder())
}
