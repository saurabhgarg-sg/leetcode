package main

import (
	"fmt"
	"strconv"
	"strings"
)

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in
// reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a
// linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	l1 := createList([]int{2, 4, 3})
	l2 := createList([]int{5, 6, 4})
	addTwoNumbers(l1, l2)

	l1 = createList([]int{0})
	l2 = createList([]int{0})
	addTwoNumbers(l1, l2)

	l1 = createList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = createList([]int{9, 9, 9, 9})
	addTwoNumbers(l1, l2)

	l1 = createList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	l2 = createList([]int{5, 6, 4})
	addTwoNumbers(l1, l2)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	sum := traverse4num(l1) + traverse4num(l2)

	sumarr := strings.Split(fmt.Sprintf("%d", sum), "")
	intarr := []int{}
	for _, val := range sumarr {
		intarr = append(intarr, str2num(val))
	}
	intarr = revIntArray(intarr)
	fmt.Println(intarr)

	return createList(intarr)
}

func traverse4num(node *ListNode) int {

	arr := []int{}
	for node != nil {
		arr = append(arr, node.Val)
		node = node.Next
	}
	arr = revIntArray(arr)

	numstr := ""
	for _, val := range arr {
		numstr = numstr + fmt.Sprintf("%d", val)
	}

	return str2num(numstr)
}

func revIntArray(arr []int) []int {

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func str2num(numstr string) int {
	num, err := strconv.Atoi(numstr)
	if err != nil {
		panic("failed to convert string to int " + numstr)
	}
	return num
}

func newNode(num int) *ListNode {

	return &ListNode{Val: num, Next: nil}
}

func createList(vals []int) *ListNode {
	head := newNode(vals[0])
	node := head

	size := len(vals)
	if size > 1 {
		for i := 1; i < size; i++ {
			node.Next = newNode(vals[i])
			node = node.Next
		}
	}

	return head
}
