package main

import "fmt"

func main() {
	l1 := ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	l2 := ListNode{5, &ListNode{9, &ListNode{4, nil}}}
	p := addTwoNumbers(&l1, &l2)
	fmt.Printf("%++v", p)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := ListNode{0, nil}
	cursor := &root
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		l1Val, l2Val := 0, 0
		if l1 != nil {
			l1Val = l1.Val
		}
		if l2 != nil {
			l2Val = l2.Val
		}
		sumVal := l1Val + l2Val + carry
		carry = sumVal / 10
		sumNode := ListNode{sumVal % 10, nil}
		cursor.Next = &sumNode
		cursor = &sumNode
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return root.Next
}
