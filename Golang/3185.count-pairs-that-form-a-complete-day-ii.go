package golang

/*
 * @lc app=leetcode id=3185 lang=golang
 *
 * [3185] Count Pairs That Form a Complete Day II
 */

// @lc code=start
func countCompleteDayPairs(hours []int) int64 {
	res := int64(0)
	counts := make(map[int]int64, 24)

	for _, hour := range hours {
		target := (24 - hour%24) % 24
		res += counts[target]

		counts[hour%24]++
	}

	return res
}

// @lc code=end
