/*
 * @lc app=leetcode id=2050 lang=golang
 *
 * [2050] Parallel Courses III
 */
package main

// @lc code=start
func minimumTime(n int, relations [][]int, time []int) int {
	preq := make([][]int, n)
	memo := make([]int, n)

	for _, r := range relations {
		preq[r[0]-1] = append(preq[r[0]-1], r[1]-1)
	}

	maxT := 0
	for i := 0; i < n; i++ {
		maxT = max2050(maxT, calcTime(i, preq, time, memo))
	}
	return maxT
}

func calcTime(course int, graph [][]int, time []int, memo []int) int {
	if memo[course] != 0 {
		return memo[course]
	}

	maxT := 0
	for _, preq := range graph[course] {
		maxT = max2050(maxT, calcTime(preq, graph, time, memo))
	}

	memo[course] = maxT + time[course]

	return memo[course]
}

// func minimumTime(n int, relations [][]int, time []int) int {
// 	preq := make([][]int, n)
// 	req := make([]int, n)
//
// 	for _, r := range relations {
// 		preq[r[0]-1] = append(preq[r[0]-1], r[1]-1)
// 		req[r[1]-1] += 1
// 	}
//
// 	queue := []int{}
// 	for i := 0; i < n; i++ {
// 		if req[i] == 0 {
// 			queue = append(queue, i)
// 		}
// 	}
//
// 	req = nil
//
// 	maxT := 0
// 	for len(queue) > 0 {
// 		prevCourse := queue[0]
// 		v := preq[prevCourse]
// 		queue = queue[1:]
//
// 		currTime := time[prevCourse] % (1 << 15)
// 		currPreqTime := time[prevCourse] >> 16
//
// 		if currTime+currPreqTime > maxT {
// 			maxT = currTime + currPreqTime
// 		}
// 		for _, nextCourse := range v {
// 			nextTime := time[nextCourse] % (1 << 15)
// 			nextPreqTime := time[nextCourse] >> 16
// 			pot := currPreqTime + currTime + nextTime
// 			if pot > nextTime+nextPreqTime {
// 				time[nextCourse] = nextTime | (currTime+currPreqTime)<<16
// 			}
// 			queue = append(queue, nextCourse)
// 		}
// 	}
// 	// for prevCourse, v := range preq {
// 	// 	for _, nextCourse := range v {
// 	// 		pot := timeCopy[prevCourse] + time[nextCourse]
// 	// 		timeCopy[nextCourse] = max2050(pot, timeCopy[nextCourse])
// 	// 	}
// 	// }
//
// 	// fmt.Println(timeCopy)
// 	return maxT
// }

func max2050(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
