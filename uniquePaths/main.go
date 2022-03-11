package main

import "fmt"

func main() {
	m := 51
	n := 7

	res := uniquePaths(m, n)

	fmt.Println(res)
}

func uniquePaths(m int, n int) int {
	if m == 2 {
		return n
	}
	if n == 2 {
		return m
	}
	if m == 1 && n == 1 {
		return 1
	}
	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}
