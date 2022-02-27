package main

import "fmt"

func main() {
	digit := "23"
	res := letterCombinations(digit)
	fmt.Printf("res : %+v", res)
}

var M = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {

	var result []string
	if len(digits) == 0 {
		return result
	}

	recursives(digits, 0, "", &result)
	return result
}

func recursives(digits string, index int, s string, result *[]string) {
	fmt.Printf("sting : %s ,round : %d", s, index)
	fmt.Println()
	if len(digits) == index {
		*result = append(*result, s)
		return
	}

	for _, c := range M[digits[index]] {
		recursives(digits, index+1, s+string(c), result)
	}

}
