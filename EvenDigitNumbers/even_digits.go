package main

import "fmt"

// Given an array nums of integers, return how many of them contain an even number of digits.
// Constraints:
// 1 <= nums.length <= 500
// 1 <= nums[i] <= 100,000

func main() {

	nums := []int{12, 345, 2, 6, 7896, 555, 901, 482, 1771}

	// As the range of values is fixed, its simpler to solve this one.
	count := 0
	for _, num := range nums {
		if num >= 10 && num <= 99 ||
			num >= 1_000 && num <= 9_999 ||
			num == 100_000 {
			count++
		}
	}
	fmt.Println(count)
}
