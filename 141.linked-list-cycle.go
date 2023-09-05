package main

type ListNode struct {
	Next *ListNode
	Val  int
}

func fastSlowPointer(head *ListNode) bool {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	const MAX int = 1000000

	for head != nil {
		if head.Val != MAX {
			head.Val = MAX
			head = head.Next
		} else {
			return true
		}
	}
	return false
}

// @lc code=end
