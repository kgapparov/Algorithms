package main

import "fmt"

/*
	Task: write the function wich gets input two arrays of integers and return true in case of each element of array1 has same element in square value on second array.
	counts of elements is matter.
	example:
	{1,2,3,4,5} and {1,4,9,25,16} - return true. order not in case
	{1,2,3,4,5} and {1,3,25,121} - false
	{1,2 } and {1,4,3} - false because diff lenth
*/
func same(a, b []int) bool {
	// make two object will count each element from each array
	var frequencyCounter1, frequencyCounter2 map[int]int
	frequencyCounter1, frequencyCounter2 = make(map[int]int), make(map[int]int)
	lena, lenb := len(a), len(b)
	if lena != lenb {
		return false
	}
	for i := 0; i < lena; i++ {
		if _, ok := frequencyCounter1[a[i]]; !ok {
			frequencyCounter1[a[i]] = 1
		} else {
			frequencyCounter1[a[i]]++
		}
	}
	for i := 0; i < lenb; i++ {
		if _, ok := frequencyCounter2[b[i]]; !ok {
			frequencyCounter2[b[i]] = 1
		} else {
			frequencyCounter2[b[i]]++
		}
	}
	//compare two objects where values same and key2 equal key1*key1
	for key, val := range frequencyCounter1 {
		if val2, ok := frequencyCounter2[key*key]; !ok || val != val2 {
			return false
		}
	}
	return true
}

func main() {
	var a, b []int
	a, b = []int{1, 2, 3, 4, 5, 5, 5, 5}, []int{1, 4, 9, 16, 25, 25, 25, 25}
	fmt.Println(same(a, b))
}
