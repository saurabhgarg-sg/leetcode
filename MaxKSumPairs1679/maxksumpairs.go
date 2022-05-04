package main

import "fmt"

// You are given an integer array nums and an integer k.
// In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.
// Return the maximum number of operations you can perform on the array.

func main() {

	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5) == 2)
	fmt.Println(maxOperations([]int{1, 2, 3, 4, 2, 3}, 5) == 3)
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6) == 1)
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 15) == 0)
}

func maxOperations(nums []int, k int) int {

	// 2 pass solution
	// First pass: create a map of values and count on how many times it occurs
	// Second pass: for each element calculate difference from total & occurences to get the op count
	elements := make(map[int]int)

	for _, val := range nums {
		if mval, exists := elements[val]; exists {
			elements[val] = mval + 1
		} else {
			elements[val] = 1
		}
	}

	ops := 0
	for _, val := range nums {
		// Remove the original element to avoid extra count.
		if mval, exists := elements[val]; exists {
			if mval == 1 {
				delete(elements, val)
			} else {
				elements[val] = mval - 1
			}

			diff := k - val
			// find the K-sum pair.
			if mval, exists := elements[diff]; exists {
				ops = ops + 1

				// Remove only one occurence for each k-sum pair.
				if mval == 1 {
					delete(elements, diff)
				} else {
					elements[diff] = mval - 1
				}
			}
		}
	}

	return ops
}
