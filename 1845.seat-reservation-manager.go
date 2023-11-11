/*
 * @lc app=leetcode id=1845 lang=golang
 *
 * [1845] Seat Reservation Manager
 */

package main

import (
	"container/heap"
)

// @lc code=start
type IntHeapSM []int

func (h IntHeapSM) Len() int           { return len(h) }
func (h IntHeapSM) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeapSM) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeapSM) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeapSM) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SeatManager struct {
	Seats   *IntHeapSM
	Counter int
}

func SMConstructor(n int) SeatManager {
	h := IntHeapSM{}
	heap.Init(&h)
	return SeatManager{&h, 1}
}

func (sm *SeatManager) Reserve() int {
	if sm.Seats.Len() != 0 {
		return heap.Pop(sm.Seats).(int)
	}

	sm.Counter++
	return sm.Counter - 1
}

func (sm *SeatManager) Unreserve(seatNumber int) {
	if seatNumber != sm.Counter-1 {
		heap.Push(sm.Seats, seatNumber)
	} else {
		sm.Counter--
	}
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
// @lc code=end
