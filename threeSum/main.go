package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-2, 0, 1, 1, 2}

	result := threeSum(nums)

	fmt.Printf("result : %+v", result)
}

func threeSum(nums []int) [][]int {

	res := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums)-1

		for j < k {

			fmt.Printf("i : %d, j %d: , k : %d", i, j, k)
			fmt.Println()

			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}

			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}

			sum := nums[i] + nums[j] + nums[k]
			fmt.Println(sum)
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return res
}
