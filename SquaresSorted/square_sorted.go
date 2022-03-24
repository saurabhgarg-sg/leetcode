package main

import (
	"fmt"
	"sort"
)

// Given an integer array nums sorted in non-decreasing order, return an array of the squares
// of each number sorted in non-decreasing order.
//Constraints:
// 1 <= nums.length <= 10,000
// -10,000 <= nums[i] <= 10,000
// nums is sorted in non-decreasing order.

func main() {
	nums := []int{-7, -3, 2, 3, 11}

	for i, num := range nums {
		nums[i] = num * num
	}
	sort.Ints(nums)
	fmt.Println(nums)
}
