package main

import (
	"errors"
	"fmt"
)

func main() {
	var a []int
	a = []int{1, 5, 4, 3, 1, 3, 123, 65}
	b, c, err := getPair(a, 22)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c, b)
	}
}

func getPair(a []int, sum int) (int, int, error) {
	if len(a) == 0 {
		err := errors.New("there no elements")
		return 0, 0, err
	}
	var tmp map[int]int
	tmp = make(map[int]int)
	length := len(a)
	for i := 0; i < length; i++ {
		if val, ok := tmp[a[i]]; !ok {
			tmp[sum-a[i]] = a[i]
		} else {
			return a[i], val, nil
		}
	}
	err := errors.New("No pairs like add up to sum")
	return 0, 0, err
}
