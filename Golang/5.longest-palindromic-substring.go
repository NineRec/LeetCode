package golang

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	// look each s-i, and expend to left + right.
	// case 1: s-i in the middle; case 2: s[i] == s[i-1]

	result := string(s[0])

	for i := 1; i < len(s); i++ {
		tmp := getLongest(s, i-1, i+1)
		if len(result) < len(tmp) {
			result = tmp
		}

		if s[i] == s[i-1] {
			tmp := getLongest(s, i-2, i+1)

			if len(result) < len(tmp) {
				result = tmp
			}
		}
	}

	return result
}

func getLongest(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return string(s[left+1 : right])
}

// @lc code=end
