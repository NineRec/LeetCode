package golang

import (
	"slices"
)

/*
 * @lc app=leetcode id=435 lang=golang
 *
 * [435] Non-overlapping Intervals
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[1] - b[1]
	})

	ans, right := 1, intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= right {
			ans++
			right = intervals[i][1]
		}
	}

	return len(intervals) - ans
}

// @lc code=end
