package main

import "fmt"

type node struct {
	val  string
	next *node
}
type singleLinkedList struct {
	head, tail *node
	length     int
}

func (sll *singleLinkedList) new() {
	sll.head, sll.tail = nil, nil
	sll.length = 0
}
func (sll *singleLinkedList) push(val string) {
	var newNode node
	newNode.new(val)
	if sll.head == nil {
		sll.head = &newNode
		sll.tail = &newNode
	} else {
		sll.tail.next = &newNode
		sll.tail = &newNode
	}
	sll.length++
}
func (sll *singleLinkedList) pop() {
	if sll.length == 0 {
		sll.head, sll.tail = nil, nil
	}
	current, newTail := sll.head, sll.head
	for current.next != nil {
		newTail = current
		current = current.next
	}
	sll.tail = newTail
	sll.tail.next = nil
	sll.length--
}
func (sll *singleLinkedList) shift() {
	if sll.length == 0 {
		return
	}
	sll.head = sll.head.next
	sll.length--
	if sll.length == 0 {
		sll.tail = nil
	}
}

func (sll *singleLinkedList) print() {
	current := sll.head
	for current != nil {
		fmt.Printf("%v\t", current.val)
		current = current.next
	}
	fmt.Println()
	fmt.Println("and total count of elements is :", sll.length)
}

func (sll *singleLinkedList) unshift(val string) {
	var newNode node
	newNode.new(val)
	if sll.length == 0 {
		sll.head = &newNode
		sll.tail = &newNode
		sll.length++
		return
	}
	newNode.next, sll.head = sll.head, &newNode
	sll.length++
}
func (sll *singleLinkedList) get(index int) *node {
	if index < 0 || index > sll.length {
		return nil
	}
	current := sll.head
	counter := 0
	for counter != index {
		current = current.next
		counter++
	}
	return current
}
func (sll *singleLinkedList) set(index int, val string) bool {
	if sll.get(index) != nil {
		sll.get(index).val = val
		return true
	}
	return false
}
func (sll *singleLinkedList) insert(index int, val string) bool {
	if index < 0 || index > sll.length-1 {
		return false
	} else if index == 0 {
		sll.unshift(val)
	} else if index == sll.length-1 {
		sll.push(val)
	} else {
		var newNode node
		var pre *node
		newNode.new(val)
		pre = sll.get(index - 1)
		newNode.next = pre.next
		pre.next = &newNode
		sll.length++
	}
	return true
}

func (sll *singleLinkedList) remove(index int) *node {
	var current *node
	if index < 0 || index > sll.length-1 {
		return nil
	} else if index == 0 {
		sll.shift()
	} else if index == sll.length-1 {
		sll.pop()
	} else {
		var prev *node
		prev = sll.get(index - 1)
		current = prev.next
		prev.next = prev.next.next
		sll.length--
	}
	return current
}
func (sll *singleLinkedList) reverse() {
	var current, prev, next *node
	sll.head = sll.tail
	current = sll.tail
	prev = nil
	for i := 0; i < sll.length; i++ {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
}
func (n *node) new(i string) {
	n.val = i
	n.next = nil
}

func main() {
	var linkedList singleLinkedList
	linkedList.new()
	linkedList.push("hi there")
	linkedList.push("good by")
	linkedList.push(" there ")
	linkedList.shift()
	linkedList.shift()
	linkedList.shift()
	linkedList.print()
	linkedList.unshift("hi there ")
	linkedList.unshift("again")
	linkedList.unshift("one more ")
	linkedList.print()
	fmt.Println(linkedList.get(1))
	linkedList.set(0, "changed value ")
	linkedList.insert(1, "inserted at 1 ")
	linkedList.insert(2, "inserted at 2")
	linkedList.print()
	linkedList.remove(2)
	linkedList.print()
	linkedList.remove(0)
	linkedList.print()
	linkedList.reverse()
	linkedList.print()
	//fmt.Println(linkedList)
}
