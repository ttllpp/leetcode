package main

import ()

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tempResult := result
	lastSum := 0
	for l1 != nil || l2 != nil || lastSum > 0 {
		sum := lastSum
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		//保留进位用于下一轮计算
		lastSum = sum / 10
		//计算当前位
		tempResult.Next = &ListNode{Val: sum % 10}
		tempResult = tempResult.Next
	}
	return result.Next
}
