package main

import "fmt"

func power(a, b int) int {
	if b == 1 {
		return a
	}
	if b == 0 {
		return 1
	} else {
		return a * power(a, b-1)
	}
}
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
func productOffArray(a []int) int {
	if len(a) == 1 {
		return a[0]
	}
	current := a[0]
	a = a[1:]
	return current * productOffArray(a)
}
func recursiveRange(n int) int {
	if n == 1 {
		return 1
	}
	return n + recursiveRange(n-1)
}
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func recursiveString(s string) string {
	var result []byte
	if len(s) == 1 {
		return string(s[0])
	}
	result = append(result, s[len(s)-1])
	result = append(result, []byte(recursiveString(s[:len(s)-1]))...)
	return string(result)
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	if len(s) == 2 {
		return s[0] == s[1]
	}
	if s[0] == s[len(s)-1] {
		return isPalindrome(s[1 : len(s)-1])
	}
	return false
}
func main() {
	fmt.Println(power(123, 3))
	fmt.Println(productOffArray([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(factorial(5))
	fmt.Println(recursiveRange(6))
	fmt.Println(fib(10))
	fmt.Println(recursiveString("Hi my name is Khassangali"))
	fmt.Println(isPalindrome("ohooho"))
}
