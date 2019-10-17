package main

import "fmt"

func bubleSort(a []int) []int {
	isChanged := true
	for i := len(a); i > 0; i-- {
		isChanged = false
		for j := 0; j < i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				isChanged = true
			}
		}
		if !isChanged {
			break
		}
	}
	return a
}

func selectionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		index := i
		for j := i; j < len(a); j++ {
			min := a[i]
			if min > a[j] {
				index = j
			}
		}
		if i != index {
			a[i], a[index] = a[index], a[i]
		}
	}
	return a
}

func insertionSort(a []int) []int {
	var j int
	for i := 1; i < len(a); i++ {
		current := a[i]
		for j = i - 1; j >= 0 && a[j] > current; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = current
	}
	return a
}

func mergeArrays(arr1, arr2 []int) []int {
	var i, j int
	var result []int
	for i < len(arr1) && j < len(arr2) {
		if arr2[j] > arr1[i] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	if i < len(arr1) {
		result = append(result, arr1[i:]...)
	}
	if j < len(arr2) {
		result = append(result, arr2[j:]...)
	}
	return result
}

func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	middle := len(a) / 2
	fmt.Println(a)
	left := mergeSort(a[:middle])
	right := mergeSort(a[middle:])
	return mergeArrays(left, right)
}

func pivot(arr []int, start int, end int) int {
	var pivot = arr[start]
	var swapIdx = start
	for i := start + 1; i < end; i++ {
		if pivot > arr[i] {
			swapIdx++
			arr[i], arr[swapIdx] = arr[swapIdx], arr[i]
		}
	}
	arr[start], arr[swapIdx] = arr[swapIdx], arr[start]
	return swapIdx
}

// func quickSort(a []int, left, right int) []int {
// 	right = len(a)
// 	if left < right {
// 		var pivotIndex = pivot(a, left, right)
// 		quickSort(a, start, pivotIndex-1)
// 		quickSort(a, pivotIndex+1)
// 	}
// 	return a
// }
func main() {
	a := []int{23, 2, 33, 23, 1, 234, 12}
	b := []int{23, 23, 22, 34, 2345, 2, 1}
	fmt.Println("We will test sorting algorithms: ")
	fmt.Println("This is input array : ")
	fmt.Println(a)
	fmt.Println("This is result of BubleSort :")
	fmt.Println(bubleSort(a))
	fmt.Println("This is result of SelectionSort ")
	fmt.Println(selectionSort(a))
	fmt.Println("This is result of InserstionSort: ")
	fmt.Println(insertionSort(a))
	fmt.Println("Testing merging function and arr1 is :", b)
	fmt.Println("Testing merging function and arr1 is :", a)
	fmt.Println("after merging result is :")
	fmt.Println(mergeArrays(bubleSort(a), bubleSort(b)))
	fmt.Println("Merge Sort results: ")
	fmt.Println(mergeSort(append(a, b...)))
}
