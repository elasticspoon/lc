/*
 * @lc app=leetcode id=1425 lang=golang
 *
 * [1425] Constrained Subsequence Sum
 */

package main

import (
	"container/heap"
)

// @lc code=start

type SlidingWindowHeap []*dpItem

type dpItem struct {
	index   int
	dpIndex int
	value   int
}

func (pq SlidingWindowHeap) Len() int { return len(pq) }
func (pq SlidingWindowHeap) Less(i, j int) bool {
	return pq[i].value > pq[j].value
}

func (pq SlidingWindowHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (wh *SlidingWindowHeap) Push(x any) {
	n := len(*wh)
	item := x.(*dpItem)
	item.index = n
	*wh = append(*wh, item)
}

func (pq *SlidingWindowHeap) Peek() *dpItem {
	return (*pq)[0]
}

// func (wp *SlidingWindowHeap) remove(item *dpItem, index int) {
// 	heap.Remove(wp, index)
// 	// wp.Push(item)
// }

func (wh *SlidingWindowHeap) Pop() any {
	old := *wh
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*wh = old[0 : n-1]
	return item
}

func constrainedSubsetSum(nums []int, k int) int {
	dp := make([]*dpItem, len(nums))
	dp[0] = &dpItem{index: 0, value: nums[0], dpIndex: 0}
	max := nums[0]

	pq := make(SlidingWindowHeap, 0)
	heap.Init(&pq)
	heap.Push(&pq, dp[0])

	for i := 1; i < len(nums); i++ {
		potential := max1425(nums[i], pq.Peek().value+nums[i])

		item := &dpItem{value: potential, dpIndex: i}

		heap.Push(&pq, item)
		if pq.Len() > k {

			last := dp[i-k]
			heap.Remove(&pq, last.index)
		}

		dp[i] = item

		if potential > max {
			max = potential
		}
	}

	return max
}

func max1425(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
