package golang

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	result := string(s[0])

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			left, right := diffDirections(s, i-1, i)
			if right-left+1 > len(result) {
				result = string(s[left : right+1])
			}
		}

		if i+1 < len(s) && s[i-1] == s[i+1] {
			left, right := diffDirections(s, i-1, i+1)
			if right-left+1 > len(result) {
				result = string(s[left : right+1])
			}
		}
	}

	return result
}

func diffDirections(s string, left, right int) (int, int) {
	for left > 0 && right < len(s)-1 && s[left-1] == s[right+1] {
		left--
		right++
	}

	return left, right
}

// @lc code=end
