package myutils

import (
	"fmt"
)

type node struct {
	val  interface{}
	next *node
}

func (n *node) new(val interface{}) {
	n.val = val
	n.next = nil
}

type Queue struct {
	first, last *node
	size        int
}

func (q *Queue) New() {
	q.first, q.last = nil, nil
	q.size = 0
}

func (q *Queue) Enqueue(n *node) int {
	if q.size == 0 {
		q.first, q.last = n, n
	} else {
		q.last.next = n
		q.last = n
	}
	q.size++
	return q.size
}

func (q *Queue) Dequeue() *node {
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

func (q *Queue) Print() {
	current := q.first
	for current != nil {
		fmt.Printf("%d ", current.val)
		current = current.next
	}
	fmt.Println()
	fmt.Println("the current size is :", q.size)
}
