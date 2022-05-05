package main

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

	l3Head := newNode(0)
	l3 := l3Head

	// Traverse lists till both are nil. One can be nil while other is not done.
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}

		sum = sum + carry
		carry = 0
		// carry will always be 1
		if sum > 9 {
			carry = 1
		}

		l3.Val = sum % 10
		if l1 != nil || l2 != nil || carry == 1 {
			// carry is always 1 so for last possible after ending the lists.
			l3.Next = newNode(1)
		}
		l3 = l3.Next
	}

	return l3Head
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
