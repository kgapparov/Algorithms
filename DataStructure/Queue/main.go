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

type queue struct {
	first, last *node
	size        int
}

func (q *queue) new() {
	q.first, q.last = nil, nil
	q.size = 0
}

func (q *queue) enqueue(str string) int {
	var newNode node
	newNode.new(str)
	if q.size == 0 {
		q.first, q.last = &newNode, &newNode
	} else {
		q.last.next = &newNode
		q.last = &newNode
	}
	q.size++
	return q.size
}
func (q *queue) dequeue() *node {
	var delNode *node
	if q.size == 0 {
		return nil
	} else if q.size == 1 {
		q.first, q.last = nil, nil
	} else {
		delNode = q.first
		q.first = q.first.next
	}
	q.size--
	return delNode
}

func (q *queue) print() {
	current := q.first
	for current != nil {
		fmt.Printf("%s\t", current.val)
		current = current.next
	}
	fmt.Println()
	fmt.Println("the current size is :", q.size)
}

func main() {
	var queue queue
	queue.new()
	queue.enqueue("100")
	queue.enqueue("200")
	queue.enqueue("300")
	queue.print()
	queue.dequeue()
	queue.print()
	queue.dequeue()
	queue.print()
	queue.dequeue()
	queue.print()
	queue.dequeue()
}
