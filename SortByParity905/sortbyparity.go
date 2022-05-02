package main

import "fmt"

func main() {
	nums := []int{3, 1, 2, 4}

	fmt.Println(sortArrayByParity(nums))

}

func sortArrayByParity(nums []int) []int {
	evenIndex := 0
	size := len(nums)
	oddIndex := size - 1
	retArray := make([]int, size)

	for _, val := range nums {

		if val%2 == 0 {
			// place even number from the begining of array
			retArray[evenIndex] = val
			evenIndex = evenIndex + 1
		} else {
			// Odd numbers go the end
			retArray[oddIndex] = val
			oddIndex = oddIndex - 1
		}
	}

	return retArray
}
