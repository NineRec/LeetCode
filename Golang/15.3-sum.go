package golang

import (
	"slices"
)

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	result := [][]int{}
	i := 0
	for i < len(nums)-2 {
		target := 0 - nums[i]

		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right] == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				left = nextDiff(nums, left)
				right--
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				left++
			}
		}

		i = nextDiff(nums, i)
	}

	return result
}

func nextDiff(nums []int, idx int) int {
	next := idx + 1
	for next < len(nums) && nums[next] == nums[idx] {
		next++
	}
	return next
}

// @lc code=end
