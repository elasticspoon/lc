/*
 * @lc app=leetcode id=1584 lang=golang
 *
 * [1584] Min Cost to Connect All Points
 */

package main

import (
	"container/heap"
	"math"
)

// @lc code=start
type Item struct {
	distance, point int
}

//	type Item struct {
//		distance, point int
//	}
type PriorityQueue1584 []Item

func (pq PriorityQueue1584) Len() int { return len(pq) }

func (pq PriorityQueue1584) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue1584) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue1584) Push(x interface{}) {
	item := x.(Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue1584) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func manhattanDistance(p1, p2 []int) int {
	return int(math.Abs(float64(p1[0]-p2[0])) + math.Abs(float64(p1[1]-p2[1])))
}

//	func minCostConnectPoints(points [][]int) int {
//		heapDict := make(map[int]int)
//		for i := range points {
//			heapDict[i] = int(math.MaxInt64) // Initialize all distances to infinity
//		}
//		heapDict[0] = 0 // Start node
//
//		queue := make(PriorityQueue, 1)
//		queue[0] = Item{distance: 0, point: 0}
//		heap.Init(&queue)
//
//		visited := make(map[int]bool, len(points))
//		mstWeight := 0
//
//		for queue.Len() > 0 {
//			current := heap.Pop(&queue).(Item)
//			weight, point := current.distance, current.point
//
//			if visited[point] || heapDict[point] < weight {
//				continue
//			}
//
//			visited[point] = true
//			mstWeight += weight
//
//			for v := range points {
//				if !visited[v] {
//
//					newDistance := manhattanDistance(points[point], points[v])
//
//					if newDistance < heapDict[v] {
//						heapDict[v] = newDistance
//						heap.Push(&queue, Item{distance: newDistance, point: v})
//					}
//				}
//			}
//		}
//
//		return mstWeight
//	}
func minCostConnectPoints(points [][]int) int {
	minDist := make(map[int]int)
	for i := range points {
		minDist[i] = int(math.MaxInt64) // Initialize all distances to infinity
	}
	minDist[0] = 0 // Start node

	queue := make(PriorityQueue1584, 1)
	queue[0] = Item{distance: 0, point: 0}
	heap.Init(&queue)

	visited := make(map[int]bool, len(points))
	sum := 0

	for queue.Len() > 0 {
		current := heap.Pop(&queue).(Item)
		currDist, point := current.distance, current.point

		if visited[point] || minDist[point] < currDist {
			continue
		}

		visited[point] = true
		sum += currDist

		for v := range points {
			if !visited[v] {
				newDistance := manhattanDistance(points[point], points[v])

				if newDistance < minDist[v] {
					minDist[v] = newDistance
					heap.Push(&queue, Item{distance: newDistance, point: v})
				}
			}
		}
	}

	return sum
}

// @lc code=end
