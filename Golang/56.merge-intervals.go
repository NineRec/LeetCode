package golang

import "sort"

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		tmp := merged[len(merged)-1]
		if intervals[i][0] <= tmp[1] {
			if intervals[i][1] > tmp[1] {
				tmp[1] = intervals[i][1]
			}
			continue
		}

		merged = append(merged, intervals[i])
	}
	return merged
}

// @lc code=end
