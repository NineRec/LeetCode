/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	maxArea := heights[0]

	for i := 0; i < len(heights); i++ {
		left, right := i, i
		for j := i - 1; j >= 0 && heights[j] >= heights[i]; j-- { // left part
			left--
		}
		for j := i + 1; j < len(heights) && heights[j] >= heights[i]; j++ { // left part
			right++
		}

		maxI := heights[i] * (right - left + 1)
		maxArea = max(maxArea, maxI)
	}

	return maxArea
}

// @lc code=end

