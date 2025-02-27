/*
 * @lc app=leetcode id=214 lang=golang
 *
 * [214] Shortest Palindrome
 */

// @lc code=start
func shortestPalindrome(s string) string {
	// find the prefix that are palindrome
	right := len(s) - 1
	for ; right > 0; right-- {
		if s[0] != s[right] {
			continue
		}

		l, r := 0, right
		for s[l] == s[r] && l < r {
			l++
			r--
		}

		if l >= r { // found
			break
		}
	}

	result := ""
	for right = right + 1; right <= len(s)-1; right++ {
		result = string(s[right]) + result
	}

	result += s
	return result
}

// @lc code=end

