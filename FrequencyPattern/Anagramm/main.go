package main

import "fmt"

func validAnagram(a, b string) bool {
	//create obects to collect values and counts of each element of strings, as in go string is slice of bytes (chars) we will careate byte map with int values
	var f1, f2 map[byte]int
	f1, f2 = make(map[byte]int), make(map[byte]int)
	//len of each variables to optimise for loop
	lena, lenb := len(a), len(b)
	//if length are not same return false
	if lena != lenb {
		return false
	}
	//for loop for each of elements and collect count of each rune in each string inputs
	for i := 0; i < lena; i++ {
		if _, ok := f1[a[i]]; !ok {
			f1[a[i]] = 1
		} else {
			f1[a[i]]++
		}
	}
	for i := 0; i < lenb; i++ {
		if _, ok := f2[b[i]]; !ok {
			f2[b[i]] = 1
		} else {
			f2[b[i]]++
		}
	}
	//now we will check if key and values are the same on each map objects
	for key, val := range f1 {
		if val2, ok := f2[key]; !ok || val != val2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(validAnagram("abcd", "cbda"))
	fmt.Println(validAnagram("aaz", "zza"))
	fmt.Println(validAnagram("anagram", "nagaram"))
	fmt.Println(validAnagram("qwerty", "qretwy"))
}
