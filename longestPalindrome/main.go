package main

import (
	"fmt"
)

func main() {

	input := "cbbd"

	Output := longestPalindrome(input)

	fmt.Printf("Input : %s", input)
	fmt.Println()
	fmt.Printf("Output : %s", Output)
	fmt.Scanln()
}

func longestPalindrome(s string) string {

	var longest int
	var longestStr string

	for i := 0; i < len(s); i++ {
		maxPalindrome(&longest, &longestStr, s, i, i)
		maxPalindrome(&longest, &longestStr, s, i, i+1)
	}

	return longestStr
}

func maxPalindrome(longest *int, longestStr *string, str string, startL int, startR int) {

	for startL >= 0 && startR < len(str) {

		if str[startL] != str[startR] {
			break
		}

		if len(str[startL:startR+1]) > *longest {
			*longest = len(str[startL : startR+1])
			*longestStr = str[startL : startR+1]
		}
		startR++
		startL--
	}
}
