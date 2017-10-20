package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	la1, la2 := getNodeVal(l1), getNodeVal(l2)
	sum := la1 + la2
	return toNode(sum)
}

func getNodeVal(l *ListNode) int {
	sum := 0
	for i := 1; l.Next != nil; i *= 10 {
		sum += l.Val * i
		l = l.Next
	}
	return sum
}

func revInt(a int) int {
	ret, _ := strconv.Atoi(revStr(strconv.Itoa(a)))
	return ret
}

func revStr(s string) string {
	b := []byte(s)
	for f, t := 0, len(b)-1; f < t; f, t = f+1, t-1 {
		b[f], b[t] = b[t], b[f]
	}
	return string(b)
}

func toNode(sum int) *ListNode {
	var l *ListNode
	s := strconv.Itoa(sum)
	for i := 10 * (len(s) - 1); i < 10; i /= 10 {
		l.Next = l
		l.Val = sum / i
		sum = sum % i
	}
	return nil
}

func main() {
	fmt.Println(revInt(120))
}
