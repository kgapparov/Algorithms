package main

import (
	"fmt"
)

func countDown(n int) {
	if n <= 0 {
		fmt.Println("All done")
		return
	}
	fmt.Println(n)
	n--
	countDown(n)
}
func power(a, b int) int {
	if b == 1 {
		return a
	}
	b--
	return a * power(a, b)
}
func sumRange(num int) int {
	if num == 1 {
		return 1
	}
	return num + sumRange(num-1)
}
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
func getOdds(a []int) []int {
	var result []int
	var getOddsHelper func(a []int)
	getOddsHelper = func(a []int) {
		//fmt.Println(a)
		if len(a) == 0 {
			return
		}
		if a[0]%2 == 0 {
			result = append(result, a[0])
		}
		a = a[1:]
		getOddsHelper(a)
	}
	getOddsHelper(a)
	return result
}

func getOddsPure(a []int) []int {
	var result []int
	if len(a) == 0 {
		return []int{}
	}
	if a[0]%2 == 0 {
		result = append(result, a[0])
	}
	a = a[1:]
	result = append(result, getOddsPure(a)...)
	return result
}

func main() {
	fmt.Println(power(5, 2))
	countDown(power(10, 1))
	fmt.Println(sumRange(3))
	fmt.Println(factorial(4))
	fmt.Println(getOddsPure([]int{1, 2, 3, 4, 5, 6}))
}
