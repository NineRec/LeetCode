package golang

import (
	"slices"
)

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	curr := 0
	merged := [][]int{intervals[curr]}
	for i := 1; i < len(intervals); i++ {
		if merged[curr][1] >= intervals[i][0] {
			merged[curr][1] = max(merged[curr][1], intervals[i][1])
		} else {
			merged = append(merged, intervals[i])
			curr++
		}
	}

	return merged
}

// @lc code=end
