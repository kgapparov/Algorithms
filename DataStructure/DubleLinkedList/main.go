package main

import "fmt"

type node struct {
	val  string
	next *node
	prev *node
}

func (n *node) new(str string) {
	n.val = str
	n.next = nil
	n.prev = nil
}

type doubleLinkedList struct {
	head, tail *node
	length     int
}

func (dll *doubleLinkedList) new() {
	dll.head, dll.tail, dll.length = nil, nil, 0
}

func (dll *doubleLinkedList) push(val string) {
	var newNode node
	newNode.new(val)
	if dll.length == 0 {
		dll.head, dll.tail = &newNode, &newNode
	} else {
		dll.tail.next = &newNode
		newNode.prev = dll.tail
		dll.tail = &newNode
	}
	dll.length++
}

func (dll *doubleLinkedList) pop() *node {
	var delNode *node
	if dll.length == 0 {
		return nil
	}
	if dll.length == 1 {
		delNode = dll.tail
		dll.head, dll.tail = nil, nil
	} else {
		delNode = dll.tail
		dll.tail = dll.tail.prev
		dll.tail.next = nil
	}
	dll.length--
	return delNode
}

func (dll *doubleLinkedList) unshift(val string) {
	var newNode node
	newNode.new(val)
	if dll.length == 0 {
		dll.head, dll.tail = &newNode, &newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = &newNode
		dll.head = &newNode
	}
	dll.length++
}

func (dll *doubleLinkedList) shift() *node {
	var delNode *node
	if dll.length == 0 {
		return nil
	}
	if dll.length == 1 {
		delNode = dll.head
		dll.head, dll.tail = nil, nil
	} else {
		delNode = dll.head
		dll.head = dll.head.next
		dll.head.prev = nil

	}
	dll.length--
	return delNode
}

func (dll *doubleLinkedList) get(index int) *node {
	if index < 0 || index >= dll.length {
		return nil
	}
	var current *node
	if index < dll.length/2 {
		counter := 0
		current = dll.head
		for counter != index {
			current = current.next
			counter++
		}
	} else {
		counter := dll.length - 1
		current = dll.tail
		for counter != index {
			current = current.prev
			counter--
		}
	}
	return current
}

func (dll *doubleLinkedList) set(index int, str string) bool {
	if dll.get(index) != nil {
		dll.get(index).val = str
		fmt.Println(dll.get(index).val)
		return true
	}
	return false
}

func (dll *doubleLinkedList) insert(index int, str string) bool {
	if index < 0 || index > dll.length {
		return false
	}
	if index == 0 {
		dll.unshift(str)
	} else if index == dll.length {
		dll.push(str)
	} else {
		var newNode node
		newNode.new(str)
		beforeNode := dll.get(index - 1)
		afterNode := beforeNode.next
		beforeNode.next = &newNode
		afterNode.prev = &newNode
		newNode.prev = beforeNode
		newNode.next = afterNode
		dll.length++
	}
	return true
}

func (dll *doubleLinkedList) remove(index int) *node {
	if index < 0 || index >= dll.length {
		return nil
	}
	if index == 0 {
		return dll.shift()
	}
	if index == dll.length-1 {
		return dll.pop()
	}
	beforeNode, afterNode := dll.get(index).prev, dll.get(index).next
	currentNode := dll.get(index)
	beforeNode.next = afterNode
	afterNode.prev = beforeNode
	currentNode.next, currentNode.prev = nil, nil
	dll.length--
	return currentNode
}
func (dll *doubleLinkedList) print() {
	current := dll.head
	for current != nil {
		fmt.Printf("%s\t", current.val)
		current = current.next
	}
	fmt.Println()
	fmt.Println("The length is: ", dll.length)
}

func main() {
	var dll doubleLinkedList
	dll.new()
	dll.push("100")
	dll.push("200")
	dll.push("300")
	dll.push("500")
	dll.push("asd;lkfj")
	dll.pop()
	dll.print()
	dll.unshift("1000000")
	dll.unshift("10000")
	dll.print()
	dll.shift()
	dll.print()
	dll.shift()
	dll.print()
	if dll.get(3) != nil {
		fmt.Println(dll.get(3).val)
	} else {
		fmt.Println("out of range ")
	}
	dll.set(2, "BLALBALBA")
	dll.insert(4, "Inserted value")
	dll.print()
	dll.remove(0)
	dll.print()
}
