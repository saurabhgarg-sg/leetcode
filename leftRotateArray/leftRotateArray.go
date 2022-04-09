package main

import "fmt"

// Given an array, left rotate the elements inline
// Do not use another array

func main() {

	arr := []int{1, 8, 2, 3, 9, 4, 5, 0}

	fmt.Println(arr)

	max := len(arr)
	firstElement := arr[0]

	for j := 0; j < max-1; j++ {
		arr[j] = arr[j+1]
	}
	arr[max-1] = firstElement
	fmt.Println(arr)
}
