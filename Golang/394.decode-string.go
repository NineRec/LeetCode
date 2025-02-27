package golang

import "strings"

/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start
func decodeString(s string) string {
	numStack := []int{1}
	strStack := []string{""}

	i := 0
	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' {
			// push in num stask
			num := 0
			for s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			numStack = append(numStack, num)
			continue
		}

		if s[i] == '[' {
			strStack = append(strStack, "")
			i++
			continue
		}

		if s[i] == ']' {
			i++

			// pop from str & num stack
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			tmp := strings.Repeat(str, num)
			strStack[len(strStack)-1] += tmp
			continue
		}

		if s[i] >= 'a' && s[i] <= 'z' {
			strStack[len(strStack)-1] += string(s[i])
			i++
		}
	}

	return strStack[0]
}

// @lc code=end
