/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func _copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// build interweaved list
	current := head
	for current != nil {
		node := Node{Val: current.Val, Next: current.Next}
		current.Next = &node
		current = current.Next.Next
	}

	// set random nodes
	current = head
	for current != nil {
		if current.Random != nil {
			current.Next.Random = current.Random.Next
		}
		current = current.Next.Next
	}

	// fix original array
	current = head
	newHead := head.Next
	for current != nil {
		newNode := current.Next
		current.Next = newNode.Next
		if newNode.Next == nil {
			newNode.Next = nil
		} else {
			newNode.Next = newNode.Next.Next
		}
		current = current.Next
	}

	return newHead
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

// this is slower
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	nodeMap := make(map[*Node]*Node)

	current := head
	for current != nil {
		nodeMap[current] = &Node{Val: current.Val}
		current = current.Next
	}

	current = head
	for current != nil {
		nodeMap[current].Random = nodeMap[current.Random]
		nodeMap[current].Next = nodeMap[current.Next]
		current = current.Next
	}

	return nodeMap[head]
}

// @lc code=end
