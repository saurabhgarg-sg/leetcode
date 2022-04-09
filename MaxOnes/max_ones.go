package main

// Given a binary array nums, return the maximum number of consecutive 1's in the array.
// Constraints:
// 1 <= nums.length <= 105
// nums[i] is either 0 or 1.

import "fmt"

func main() {

	nums := []int{1, 1, 0, 1, 1, 1}

	maxOnes := 0
	currentOnes := 0
	maxIndex := len(nums) - 1

	for index, num := range nums {

		if num == 1 {
			currentOnes = currentOnes + 1
		}

		// When last element is 1, we make decision by index.
		if num != 1 || index == maxIndex {
			if currentOnes > maxOnes {
				maxOnes = currentOnes
			}
			currentOnes = 0
		}
	}

	fmt.Println(maxOnes)
}
