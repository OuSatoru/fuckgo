package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	ret := new(ListNode)
	ret.Val = 0
	for l1 != nil && l2 != nil {
		v1 := 0
		v2 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		carry = (v1 + v2 + carry) / 10
		value := (v1 + v2 + carry) % 10
	}
}

func main() {

}
