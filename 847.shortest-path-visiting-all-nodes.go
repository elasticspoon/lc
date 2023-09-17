/*
 * @lc app=leetcode id=847 lang=golang
 *
 * [847] Shortest Path Visiting All Nodes
 */

package main

// @lc code=start
//

type NodeState struct {
	value      int
	visitedSet int
	stepCount  int
}

func queuePop(queue *[]NodeState) NodeState {
	old := *queue
	item := old[0]
	*queue = old[1:]
	return item
}

func hashState(mask, index int) int {
	return mask<<4 + index
}

func shortestPathLength(graph [][]int) int {
	visitedStates := make(map[int]bool, 0)
	queue := []NodeState{}
	target := (1 << len(graph)) - 1

	for i := range graph {
		mask := 1 << i
		queue = append(queue, NodeState{value: i, stepCount: 0, visitedSet: mask})
		// visitedStates[hashState(mask, i)] = true
	}

	// printQueue(queue)
	for len(queue) > 0 {
		current := queuePop(&queue)

		if current.visitedSet == target {
			return current.stepCount
		}

		for _, neighbor := range graph[current.value] {
			newVisted := current.visitedSet | (1 << neighbor)
			newHash := hashState(newVisted, neighbor)

			potentialState := NodeState{
				stepCount:  current.stepCount + 1,
				visitedSet: newVisted, value: neighbor,
			}

			if !visitedStates[newHash] {
				visitedStates[newHash] = true
				queue = append(queue, potentialState)
			}
		}

	}

	return -1
}

// @lc code=end
// func printQueue(q []NodeState) {
// 	for _, v := range q {
// 		fmt.Printf("Node: %d [%03b] %d | ", v.value, v.visitedSet, v.stepCount)
// 	}
// 	fmt.Println("")
// }

func main847() {
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

	// graph := [][]int{{1, 2, 3}, {0}, {0}, {0}}
	// fmt.Println(shortestPathLength(graph))
	// graph = [][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}}
	// fmt.Println(shortestPathLength(graph))
}
