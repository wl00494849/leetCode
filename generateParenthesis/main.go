package main

import "fmt"

func main() {
	res := generateParenthesis(3)
	fmt.Printf("%+v", res)
}

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	backtrack(n, n, &result, "")
	return result
}

func backtrack(left, right int, result *[]string, str string) {

	if left == 0 && right == 0 {
		*result = append(*result, str)
		return
	}

	if right < left {
		return
	}

	if left > 0 {
		backtrack(left-1, right, result, str+"(")
	}

	if right > 0 {
		backtrack(left, right-1, result, str+")")
	}
}
