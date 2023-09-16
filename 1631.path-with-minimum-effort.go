/*
 * @lc app=leetcode id=1631 lang=golang
 *
 * [1631] Path With Minimum Effort
 */
package main

import (
	"container/heap"
	"fmt"
	"math"
)

// @lc code=start
type Cell struct {
	x, y int
	dist int
}

type PriorityQueue []Cell

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(Cell)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func max1631(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minimumEffortPath(heights [][]int) int {
	maxX := len(heights)
	maxY := len(heights[0])

	dist := make([][]int, maxX)
	for i := 0; i < maxX; i++ {
		dist[i] = make([]int, maxY)
		for j := 0; j < maxY; j++ {
			dist[i][j] = math.MaxInt32
		}
	}

	dist[0][0] = 0
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	queue := &PriorityQueue{Cell{0, 0, 0}}
	heap.Init(queue)

	for queue.Len() > 0 {

		current := heap.Pop(queue).(Cell)
		cDist, x, y := current.dist, current.x, current.y

		if cDist > dist[x][y] {
			continue
		}
		if x == maxX-1 && y == maxY-1 {
			return cDist
		}

		for _, v := range directions {
			newX, newY := x+v[0], y+v[1]
			if 0 <= newX && newX < maxX && 0 <= newY && newY < maxY {
				distCToN := max1631(cDist, absDiff(heights[x][y], heights[newX][newY]))
				if distCToN < dist[newX][newY] {
					heap.Push(queue, Cell{x: newX, y: newY, dist: distCToN})
					dist[newX][newY] = distCToN
				}
			}
		}
	}

	return -1
}

// @lc code=end

func main1631() {
	// fmt.Printf("loop: (%d, %d)\n", x, y)
	//
	// for i, v := range dist {
	// 	for j, vv := range v {
	// 		if vv == math.MaxInt32 {
	// 			fmt.Printf("(%d, %d): %d | ", i, j, -1)
	// 		} else {
	// 			fmt.Printf("(%d, %d): %d | ", i, j, vv)
	// 		}
	// 	}
	// 	fmt.Printf("\n")
	// }

	heights := [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
	fmt.Println(minimumEffortPath(heights))
	heights = [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}
	fmt.Println(minimumEffortPath(heights))
	heights = [][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}
	fmt.Println(minimumEffortPath(heights))
}
