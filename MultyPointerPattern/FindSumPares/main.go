package main

import "fmt"

func sum(a []int, sum int) (int, int) {
	var left, right = 0, len(a) - 1
	for left < right {
		subsum := a[left] + a[right]
		if subsum == sum {
			return a[left], a[right]
		} else if subsum > sum {
			right--
		} else {
			left++
		}
	}
	return 0, 0
}

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4, 5, 6, 7}, 9))
}
