/*
 * @lc app=leetcode id=799 lang=golang
 *
 * [799] Champagne Tower
 */
package main

// import "fmt"

// @lc code=start
func champagneTower(poured int, query_row int, query_glass int) float64 {
	memo := make(map[int]float64, 0)
	res := champagneTowerRec(&memo, poured, query_row, query_glass)
	// fmt.Println(memo)

	if res > 1 {
		return 1
	} else if res < 0 {
		return 0
	} else {
		return res
	}
}

func champagneTowerRec(m *map[int]float64, poured int, row int, glass int) float64 {
	if glass < 0 || row < glass {
		return 0
	}

	memo := *m
	location := hashLoc(row, glass)

	if value, ok := memo[location]; ok {
		return value
	}

	if row == 0 {
		memo[location] = float64(poured)
		return memo[location]
	} else {
		left := max799((champagneTowerRec(m, poured, row-1, glass-1)-1)/2, 0)
		right := max799((champagneTowerRec(m, poured, row-1, glass)-1)/2, 0)
		result := left + right
		memo[location] = result
		return result
	}
}

func hashLoc(row int, glass int) int {
	if glass >= (row+1)/2 {
		glass = row - glass
	}
	return row<<7 | glass
}

func max799(a float64, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
// func main() {
// 	tests := []struct {
// 		poured   int
// 		row      int
// 		glass    int
// 		expected float64
// 	}{
// 		{0, 0, 0, 0},
// 		{2, 1, 1, .5},
// 		{100000009, 33, 17, 1},
// 		{4, 2, 1, .5},
// 		{6, 2, 0, .75},
// 		{6, 2, 1, 1},
// 		{25, 6, 1, 0.18750},
// 	}
//
// 	for _, test := range tests {
// 		actual := champagneTower(test.poured, test.row, test.glass)
// 		if actual != test.expected {
// 			fmt.Printf("champagneTower(%d, %d, %d) = %f, expected %f\n", test.poured, test.row, test.glass, actual, test.expected)
// 		}
// 	}
// }
