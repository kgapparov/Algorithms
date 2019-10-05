package main

import "fmt"

func countUniqueValue(a []int) int {
	var i int
	i = 0
	lena := len(a)
	if lena == 0 {
		return 0
	}
	for j := 1; j < lena; j++ {
		if a[i] != a[j] {
			i++
			a[i] = a[j]
		}
	}
	return i + 1
}
func main() {
	fmt.Println(countUniqueValue([]int{1, 1, 2, 3, 4, 5, 5, 5, 5, 6, 7}))
}
