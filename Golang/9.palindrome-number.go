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
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x = x / 10
	}
	return x == rev || x == rev/10
}

// more fast
func isPalindrome_string(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	str := strconv.FormatInt(int64(x), 10)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

// @lc code=end

