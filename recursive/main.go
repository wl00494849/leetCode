package main

import "fmt"

func main() {
	arr := [][]int{
		{3, 8},
		{9, 13},
		{4, 9},
	}

	result := recursive(arr, 0, 0, 0)

	fmt.Println(result)
	fmt.Scanln()
}

func recursive(arr [][]int, total int, row int, column int) int {

	maxColums := len(arr[0])
	maxRow := len(arr)

	fmt.Printf("maxColums : %d ", maxColums)
	fmt.Printf("maxRow : %d", maxRow)
	fmt.Println()

	fmt.Printf("row : %d", row)
	fmt.Println()
	fmt.Printf("col : %d", column)
	fmt.Println()

	total = total + arr[row][column]

	fmt.Printf("total : %d", total)
	fmt.Println()

	if row == maxRow-1 && column == maxColums-1 {
		return total
	}

	if column < maxColums-1 {
		column++
	} else {
		row++
		column = 0
	}

	return recursive(arr, total, row, column)
}
