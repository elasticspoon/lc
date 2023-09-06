/*
 * @lc app=leetcode id=725 lang=golang
 *
 * [725] Split Linked List in Parts
 */
package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func splitListToParts(head *ListNode, k int) []*ListNode {
	list := make([]*ListNode, k)

	var listLen int
	for current := head; current != nil; current = current.Next {
		listLen++
	}

	listWPart := listLen / k
	listRPart := listLen % k

	// you can do it like this but that is way ugly imo
	// for current, prev, i := head, head, 0; current != nil; prev.Next, prev, i = nil, current, i+1 {
	for current, prev, i := head, head, 0; current != nil; {
		list[i] = current

		for j := 0; j < listWPart; j++ {
			prev = current
			current = current.Next
		}

		if listRPart > 0 {
			prev = current
			current = current.Next
			listRPart--
		}

		prev.Next = nil
		prev = current
		i++
	}

	return list
}

// @lc code=end
