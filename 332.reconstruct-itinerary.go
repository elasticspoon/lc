/*
 * @lc app=leetcode id=332 lang=golang
 *
 * [332] Reconstruct Itinerary
 */

package main

import (
	"sort"
)

// @lc code=start

func findItinerary(tickets [][]string) []string {
	hash := make(map[string][]string)

	for _, v := range tickets {
		hash[v[0]] = append(hash[v[0]], v[1])
	}

	for _, v := range hash {
		sort.Slice(v, func(i, j int) bool {
			return v[i] > v[j]
		})
	}

	iten := []string{}
	stack := []string{"JFK"}

	for len(stack) > 0 {
		for len(hash[stack[len(stack)-1]]) > 0 {
			last := len(hash[stack[len(stack)-1]]) - 1
			stack = append(stack, hash[stack[len(stack)-1]][last])
			hash[stack[len(stack)-2]] = hash[stack[len(stack)-2]][:last]
		}

		iten = append(iten, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	for i := 0; i < len(iten)/2; i++ {
		iten[i], iten[len(iten)-1-i] = iten[len(iten)-1-i], iten[i]
	}

	return iten
}

// func findItinerary(tickets [][]string) []string {
// 	hash := make(map[string][]string)
//
// 	for _, v := range tickets {
// 		// start, end := v[0], v[1]
// 		// hash[start] = append(hash[start], end)
// 		hash[v[0]] = append(hash[v[0]], v[1])
// 	}
//
// 	for _, v := range hash {
// 		sort.Slice(v, func(i, j int) bool {
// 			return v[i] < v[j]
// 		})
// 	}
//
// 	iten := []string{}
//
// 	var dfs func(string)
//
// 	dfs = func(airport string) {
// 		for len(hash[airport]) > 0 {
// 			var next string
// 			next, hash[airport] = hash[airport][0], hash[airport][1:]
// 			dfs(next)
// 		}
// 		iten = append(iten, airport)
// 	}
//
// 	dfs("JFK")
//
// 	for i := 0; i < len(iten)/2; i++ {
// 		iten[i], iten[len(iten)-1-i] = iten[len(iten)-1-i], iten[i]
// 	}
//
// 	return iten
// }

// @lc code=end
