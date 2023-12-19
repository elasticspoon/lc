/*
 * @lc app=leetcode id=661 lang=golang
 *
 * [661] Image Smoother
 */

package main

// @lc code=start
func imageSmoother(img [][]int) [][]int {
	points := [][]int{{0, 0}, {0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	result := make([][]int, len(img))

	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[0]); j++ {
			sum, count := 0, 0
			for _, point := range points {
				if v, ok := getValue(img, i+point[0], j+point[1]); ok {
					sum += v
					count++
				}
			}
			result[i] = append(result[i], sum/count)
		}
	}
	return result
}

func getValue(img [][]int, i, j int) (int, bool) {
	if i < 0 || i >= len(img) || j < 0 || j >= len(img[0]) {
		return 0, false
	}
	return img[i][j], true
}

// @lc code=end
