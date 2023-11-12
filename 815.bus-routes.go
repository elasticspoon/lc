/*
 * @lc app=leetcode id=815 lang=golang
 *
 * [815] Bus Routes
 */

package main

// @lc code=start

func numBusesToDestination(routes [][]int, source int, target int) int {
	graph := make(map[int][]int)
	for i, route := range routes {
		for _, j := range route {
			graph[j] = append(graph[j], i)
		}
	}

	visited := make(map[int]bool)
	queue := [][]int{{source, 0}}
	visited[source] = true

	for len(queue) > 0 {
		stop, buses := queue[0][0], queue[0][1]
		queue = queue[1:]

		if stop == target {
			return buses
		}

		for _, i := range graph[stop] {
			for _, j := range routes[i] {
				if !visited[j] {
					queue = append(queue, []int{j, buses + 1})
					visited[j] = true
				}
			}
			routes[i] = []int{}
		}

	}
	return -1
}

// @lc code=end
