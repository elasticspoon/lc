/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

package main

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left-right == 0 {
		return head
	}

	dummy := &ListNode{Next: head}
	swapHead := dummy

	index := 1

	for ; index < left; index++ {
		swapHead = swapHead.Next
	}

	tail := swapHead.Next
	for ; index < right; index++ {
		tmp := swapHead.Next
		swapHead.Next = tail.Next
		tail.Next = tail.Next.Next
		swapHead.Next.Next = tmp
	}

	return dummy.Next
}

// func reverseBetween(head *ListNode, left int, right int) *ListNode {
// 	if head == nil || left-right == 0 {
// 		return head
// 	}
//
// 	var prev, current, next *ListNode = nil, head, head.Next
//
// 	index := 1
//
// 	for ; index < left; index++ {
// 		prev = current
// 		current = next
// 		if next != nil {
// 			next = next.Next
// 		}
// 	}
//
// 	swapStart := prev
// 	swapHead := current
//
// 	for ; index <= right; index++ {
//
// 		current.Next = prev
// 		prev = current
// 		current = next
// 		if next != nil {
// 			next = next.Next
// 		}
// 	}
//
// 	// debugHead(swap_start)
//
// 	swapHead.Next = current
//
// 	if swapStart == nil {
// 		return prev
// 	} else {
// 		swapStart.Next = prev
// 		return head
// 	}
// }

// @lc code=end
