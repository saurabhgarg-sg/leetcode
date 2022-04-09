package main

import "fmt"

// Given a fixed-length integer array arr, duplicate each occurrence of zero,
// shifting the remaining elements to the right.
// Constraints:
// 1 <= arr.length <= 10,000
// 0 <= arr[i] <= 9

func main() {

	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	maxIndex := len(arr) - 1

	for index := 0; index <= maxIndex; index++ {
		if arr[index] == 0 {
			for i := maxIndex; i > index; i-- {
				arr[i] = arr[i-1]
			}

			if index < maxIndex {
				arr[index+1] = 0
			}
			index = index + 1
		}
	}

	fmt.Println(arr)

}
