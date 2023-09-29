/*
 * @lc app=leetcode id=905 lang=golang
 *
 * [905] Sort Array By Parity
 */

package main

// @lc code=start
func sortArrayByParity(nums []int) []int {
	// fmt.Printf("Input: %v\n", nums)
	evenIndex, oddIndex := 0, len(nums)-1
	for evenIndex < oddIndex {

		if nums[evenIndex]%2 == 0 {
			evenIndex++
		}
		if nums[oddIndex]%2 != 0 {
			oddIndex--
		}
		if evenIndex >= oddIndex {
			break
		}
		if nums[evenIndex]%2 != 0 && nums[oddIndex]%2 == 0 {
			nums[evenIndex], nums[oddIndex] = nums[oddIndex], nums[evenIndex]
			evenIndex++
			oddIndex--
		}
	}
	return nums
}

// @lc code=end
// func main() {
// 	nums := []int{3, 1, 2, 4}
// 	sliceEvenAtStart(sortArrayByParity(nums))
// 	nums = []int{3, 1, 2, 4, 5, 6, 7, 8}
// 	sliceEvenAtStart(sortArrayByParity(nums))
// 	nums = []int{0}
// 	sliceEvenAtStart(sortArrayByParity(nums))
// 	nums = []int{0, 1}
// 	sliceEvenAtStart(sortArrayByParity(nums))
// }
//
// func sliceEvenAtStart(a []int) {
// 	isEven := true
// 	for _, v := range a {
// 		if v%2 != 0 && isEven {
// 			isEven = false
// 		} else if v%2 != 0 && !isEven {
// 			continue
// 		} else if v%2 == 0 && isEven {
// 			continue
// 		} else {
// 			fmt.Printf("Error: %v was expected to be even in array %v\n", v, a)
// 		}
// 	}
// }
