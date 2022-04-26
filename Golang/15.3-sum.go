package golang

import "sort"

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}
	for idx := 0; idx < len(nums)-2; idx++ {
		left := idx + 1
		right := len(nums) - 1

		for left != right {
			sum := nums[idx] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[idx], nums[left], nums[right]})
				left++
			} else if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			}
		}
	}
	return result
}

// @lc code=end
