/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

import (
	"strconv"
)

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	strX := strconv.FormatInt(int64(x), 10)
	for i := 0; i < len(strX)/2; i++ {
		if strX[i] != strX[len(strX)-i-1] {
			return false
		}
	}
	return true
}

// @lc code=end

