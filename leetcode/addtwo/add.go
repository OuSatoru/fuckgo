package main

import (
	"fmt"
	"math"
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
	l := new(ListNode)
	walk := l
	s := strconv.Itoa(sum)
	for i := int(math.Pow(10.0, float64(len(s)-1))); i >= 1; i /= 10 {
		// fmt.Println(i)
		walk.Val = sum / i
		walk.Next = new(ListNode)
		walk = walk.Next
		sum = sum % i
	}
	return l
}

func main() {
	n := toNode(1230)
	for n.Next != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
	fmt.Println(getNodeVal(n))
}
