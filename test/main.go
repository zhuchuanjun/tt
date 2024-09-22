package main

import "fmt"

func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(s); i++ {
		count += expend(s, i, i)
		count += expend(s, i, i+1)
	}
	return count
}

func expend(s string, left int, right int) int {
	count := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}

func main() {
	s := "babad"
	fmt.Println(countSubstrings(s))
}
