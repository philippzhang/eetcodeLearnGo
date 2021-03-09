package lc0002

import (
	"leetcode/structures"
)

type ListNode = structures.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	ret := &ListNode{}
	var s int = 0
	var r int = 0
	temp := ret

	for l1 != nil || l2 != nil {
		var x int = 0
		if l1 != nil {
			x = l1.Val
		}
		var y int = 0
		if l2 != nil {
			y = l2.Val
		}
		r = x + y + s
		s = r / 10
		r = r % 10
		temp.Next = &ListNode{Val: r}
		temp = temp.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if s > 0 {
		temp.Next = &ListNode{Val: 1}
	}

	return ret.Next

}
