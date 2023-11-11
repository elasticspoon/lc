package main

import "testing"

func TestDesignGraphWithShortestPathCalculator(t *testing.T) {
	graph := Constructor(4, [][]int{{0, 2, 5}, {0, 1, 2}, {1, 2, 1}, {3, 0, 3}})

	if graph.ShortestPath(3, 2) != 6 {
		t.Errorf("ShortestPath(3, 2) should be 6")
	}
	if graph.ShortestPath(0, 3) != -1 {
		t.Errorf("ShortestPath(0, 3) should be -1")
	}

	graph.AddEdge([]int{1, 3, 4})

	if graph.ShortestPath(0, 3) != 6 {
		t.Errorf("ShortestPath(0, 3) should be 6")
	}
}
