package main

import "fmt"

type node struct {
	val  string
	next *node
}

func (n *node) new(val string) {
	n.val = val
	n.next = nil
}

type stack struct {
	first, last *node
	size        int
}

func (s *stack) new() {
	s.first, s.last = nil, nil
	s.size = 0
}

func (s *stack) push(val string) int {
	var newNode node
	newNode.new(val)
	if s.size == 0 {
		s.first, s.last = &newNode, &newNode
	} else {
		newNode.next = s.first
		s.first = &newNode
	}
	s.size++
	return s.size
}
func (s *stack) pop() *node {
	var delNode *node
	if s.size == 0 {
		return nil
	}
	if s.size == 1 {
		delNode = s.first
		s.first, s.last = nil, nil
	} else {
		delNode = s.first
		s.first = s.first.next
	}
	s.size--
	return delNode
}

func (s *stack) print() {
	current := s.first
	for current != nil {
		fmt.Printf("%s\t", current.val)
		current = current.next
	}
	fmt.Println()
	fmt.Println("the current size is :", s.size)
}

func main() {
	var stack stack
	stack.new()
	stack.push("100")
	stack.push("200")
	stack.push("300")
	fmt.Println(stack.pop().val)
	stack.print()
}
