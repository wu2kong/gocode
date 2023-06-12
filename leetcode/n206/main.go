package n206

// 反转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var tail *ListNode = head
	for tail != nil {
		t := tail.Next
		tail.Next = prev
		// 往后行进
		prev = tail
		tail = t
	}
	return tail
}

func reverseList2(head *ListNode) *ListNode {
	var behind *ListNode
	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next
	}
	return behind
}
