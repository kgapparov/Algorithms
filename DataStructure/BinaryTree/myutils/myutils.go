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
	First, Last *node
	Size        int
}

func (q *Queue) New() {
	q.First, q.Last = nil, nil
	q.Size = 0
}

func (q *Queue) Enqueue(n interface{}) int {
	var newNode node
	newNode.new(n)
	if q.Size == 0 {
		q.First, q.Last = &newNode, &newNode
	} else {
		q.Last.next = &newNode
		q.Last = &newNode
	}
	q.Size++
	return q.Size
}

func (q *Queue) Dequeue() interface{} {
	var delNode node
	if q.Size == 0 {
		return nil
	} else if q.Size == 1 {
		delNode = *q.First
		q.First, q.Last = nil, nil
	} else {
		delNode = *q.First
		q.First = q.First.next
	}
	q.Size--
	return delNode.val
}

func (q *Queue) Print() {
	current := q.First
	for current != nil {
		fmt.Printf("%v ", current.val)
		current = current.next
	}
	fmt.Println()
	fmt.Println("the current Size is :", q.Size)
}
