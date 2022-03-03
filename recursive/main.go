package main

import "fmt"

func main() {
	arr := [][]int{
		{3, 8, 9},
		{9, 13},
		{4, 9, 5},
	}

	result := recursive(arr, 0, 0, 0)

	fmt.Println(result)
	fmt.Scanln()
}

func recursive(arr [][]int, total int, row int, column int) int {

	maxColums := len(arr[row]) - 1
	maxRow := len(arr) - 1

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

	if row == maxRow && column == maxColums {
		return total
	}

	if column < maxColums {
		column++
	} else {
		row++
		column = 0
	}

	return recursive(arr, total, row, column)
}
