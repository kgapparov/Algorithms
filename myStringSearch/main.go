package main

import "fmt"

func getPattern(s string) []int {
	var i, j = 1, 0
	var r = []int{0}
	var isBreak bool
	for i = 1; i < len(s); i++ {
		if s[j] == s[i] {
			r = append(r, j+1)
			j++
		} else {
			isBreak = false
			for j > 0 {
				j = r[j-1]
				if s[j] == s[i] && j > 0 {
					r = append(r, j+1)
					isBreak = true
					break
				}
			}
			if isBreak {
				break
			}
			if s[j] == s[i] {
				r = append(r, j+1)
			} else {
				r = append(r, j)
			}
		}
	}
	return r
}

func indexOf(s, subs string) int {
	var i, j = 0, 0
	var isBreak bool
	pattern := getPattern(subs)
	for i = 0; i < len(s); i++ {
		if j == len(subs)-1 {
			return i - j
		}
		if s[i] == subs[j] {
			j++
			continue
		} else {
			isBreak = false
			for j > 0 {
				j = pattern[j-1]
				if s[i] == subs[j] {
					isBreak = true
					j++
					break
				}
			}
			if isBreak {
				continue
			}

		}
	}
	return -1
}

func getCount(s, sub string) int {
	var count int
	for indexOf(s, sub) > -1 {
		s = s[indexOf(s, sub)+len(sub):]
		count++
	}
	return count
}
func main() {
	fmt.Println(getCount("testtestttestfestifaslsdfestsesttest", "est"))
}
