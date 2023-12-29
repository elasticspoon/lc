/*
 * @lc app=leetcode id=2850 lang=golang
 *
 * [2850] Minimum Moves to Spread Stones Over Grid
 */

package main

import "fmt"

// @lc code=start
func minimumMoves(grid [][]int) int {
	sum := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == 0 {
				visited := []int{}
				val, point := recDist(grid, []int{i, j}, visited)
				grid[point[0]][point[1]]--
				grid[i][j]++
				sum += val
				fmt.Println(grid)
			}
		}
	}

	return sum
}

var dirs = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func isValid(grid [][]int, point [2]int) bool {
	x, y, l := point[0], point[1], len(grid)
	return x >= 0 && x < l && y >= 0 && y < l
}

func recDist(grid [][]int, point []int, visited []int) (int, []int) {
	x, y := point[0], point[1]
	index := x*len(grid) + y

	newVisited := make([]int, len(visited))
	copy(newVisited, visited)
	newVisited = append(newVisited, index)

	if grid[x][y] > 1 {
		return 0, point
	}

	best := 999
	var bestP []int

	for _, dir := range dirs {
		newX, newY := x+dir[0], y+dir[1]
		if !isValid(grid, [2]int{newX, newY}) {
			continue
		}

		newIdx := newX*len(grid) + newY

		if containsV(newVisited, newIdx) {
			continue
		}

		v, p := recDist(grid, []int{newX, newY}, newVisited)
		v++
		if v < best {
			best = v
			bestP = p
		}
	}

	return best, bestP
}

func containsV(visited []int, idx int) bool {
	for _, v := range visited {
		if v == idx {
			return true
		}
	}
	return false
}

// @lc code=end
