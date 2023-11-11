/*
 * @lc app=leetcode id=2642 lang=golang
 *
 * [2642] Design Graph With Shortest Path Calculator
 */
package main

import (
	"container/heap"
	"math"
)

// @lc code=start
type Graph struct {
	AdjacencyList [][]*Edge
}

type Edge struct {
	Value  int
	Weight int
}

type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Edge)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Constructor(n int, edges [][]int) Graph {
	graph := Graph{AdjacencyList: make([][]*Edge, n)}
	for _, v := range edges {
		graph.AddEdge(v)
	}

	return graph
}

func (g *Graph) AddEdge(edge []int) {
	node := &Edge{Value: edge[1], Weight: edge[2]}
	g.AdjacencyList[edge[0]] = append(g.AdjacencyList[edge[0]], node)
}

func (g *Graph) ShortestPath(node1 int, node2 int) int {
	dist := make([]int, len(g.AdjacencyList))
	visited := make([]bool, len(g.AdjacencyList))
	for i := 0; i < len(dist); i++ {
		dist[i] = math.MaxInt32
	}
	dist[node1] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	pq.Push(&Edge{Value: node1, Weight: 0})

	for pq.Len() > 0 {
		edge := heap.Pop(&pq).(*Edge)
		if visited[edge.Value] {
			continue
		}
		if edge.Value == node2 {
			return dist[edge.Value]
		}

		visited[edge.Value] = true
		for _, v := range g.AdjacencyList[edge.Value] {
			potential := dist[edge.Value] + v.Weight
			if !visited[v.Value] && dist[v.Value] > potential {
				dist[v.Value] = dist[edge.Value] + v.Weight
				heap.Push(&pq, &Edge{Value: v.Value, Weight: dist[v.Value]})
			}
		}
	}

	return -1
}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */
// @lc code=end
