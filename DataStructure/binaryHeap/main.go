package main

import "fmt"

type maxBinaryHeap struct {
	value []int
}

func (m *maxBinaryHeap) Insert(value int) {
	m.value = append(m.value, value)
	m.bubleUp()
}
func (m *maxBinaryHeap) bubleUp() {
	var idx = len(m.value) - 1
	element := m.value[idx]
	for idx > 0 {
		parentIndex := (idx - 1) / 2
		parent := m.value[parentIndex]
		if element <= parent {
			break
		}
		m.value[idx], m.value[parentIndex] = m.value[parentIndex], m.value[idx]
		idx = parentIndex
	}
}

func (m *maxBinaryHeap) Extract() int {
	max := m.value[0]
	end := m.value[len(m.value)-1]
	m.value = m.value[:len(m.value)-1]
	if len(m.value) > 0 {
		m.value[0] = end
		m.sinkDown()
	}
	return max
}
func (m *maxBinaryHeap) sinkDown() {
	var idx = 0
	length := len(m.value)
	element := m.value[0]
	for {
		leftChildIdx := 2*idx + 1
		rightChildIdx := 2*idx + 1
		var leftChild, rightChild int
		var swap interface{} = nil
		if leftChildIdx < length {
			leftChild = m.value[leftChildIdx]
			if leftChild > element {
				swap = leftChildIdx
				//	fmt.Println(swap)
			}
		}
		if rightChildIdx < length {
			rightChild = m.value[rightChildIdx]
			if (rightChild > element) && (swap == nil) ||
				(swap != nil) && (rightChild > leftChild) {
				swap = rightChildIdx
			}
		}
		if swap == nil {
			//	fmt.Println("nill")
			break
		}
		m.value[idx] = m.value[swap.(int)]
		m.value[swap.(int)] = element
		idx = swap.(int)
	}
}

func main() {
	var maxBH maxBinaryHeap
	for i := 0; i < 16; i++ {
		maxBH.Insert(i)
	}
	maxBH.Insert(41)
	maxBH.Insert(39)
	maxBH.Insert(33)
	maxBH.Insert(18)
	maxBH.Insert(27)
	maxBH.Insert(12)
	maxBH.Insert(55)
	fmt.Println(maxBH.value)
	fmt.Println(maxBH.Extract())
	fmt.Println(maxBH.value)
}
