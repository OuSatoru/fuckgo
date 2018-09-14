package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	ret := new(ListNode)
	tmp := ret
	for l1 != nil || l2 != nil || carry > 0 {
		s := carry
		if l1 != nil {
			s += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			s += l2.Val
			l2 = l2.Next
		}
		carry = s / 10
		value := s % 10
		tmp.Next = &ListNode{Val: value}
		tmp = tmp.Next
	}
	return ret.Next
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	res := addTwoNumbers(l1, l2)
	fmt.Println(*res)
}
