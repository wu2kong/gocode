package n160

// 相交链表，使用链表长度拼接相等的长度值

import "fmt"

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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//boundary check
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	//if a & b have different len, then we will stop the loop after second iteration
	for a != b {
		//for the end of first iteration, we just reset the pointer to the head of another linkedlist
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
		fmt.Printf("a = %v b = %v\n", a, b)
	}
	return a
}
