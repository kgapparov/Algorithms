package main

import "fmt"

func maxSum(a []int, n int) int {
	lena := len(a)
	if lena < n {
		return 0
	}
	var maxSum, temp int
	for i := 0; i < n; i++ {
		maxSum += a[i]
	}
	temp = maxSum
	for i := n; i < lena; i++ {
		temp = temp - a[i-n] + a[i]
		if maxSum < temp {
			maxSum = temp
		}
	}
	return maxSum
}
func main() {
	fmt.Println(maxSum([]int{1, 2, 3, 4, 5, 4, 4, 4, 6, 2, 3, 4, 23, 4, 4, 23, 2}, 5))
}
