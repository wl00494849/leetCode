package main

import (
	"fmt"
)

func main() {

	input := "abcabcbb"
	res := lengthOfLongestSubstring(input)

	fmt.Printf("input string %s", input)
	fmt.Println()
	fmt.Printf("Output %d ", res)

	fmt.Scanln()
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	temp := map[byte]int{}
	var start int
	var longest int

	for i := 0; i < len(s); i++ {

		t := s[i]

		if _, ok := temp[t]; ok && temp[t] >= start {
			start = temp[t] + 1
		}

		longest = max(longest, i-start+1)
		temp[t] = i
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
