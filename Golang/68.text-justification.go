/*
 * @lc app=leetcode id=68 lang=golang
 *
 * [68] Text Justification
 */

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	res := []string{}

	left := 0
	currLen := 0
	for idx, word := range words {
		if currLen+len(word)+(idx-left) > maxWidth {
			tmp := appendLine(words[left:idx], currLen, maxWidth)
			res = append(res, tmp)

			currLen = len(word)
			left = idx
		} else {
			currLen += len(word)
		}
	}

	lastLine := strings.Join(words[left:], " ")
	res = append(res, lastLine+strings.Repeat(" ", maxWidth-len(lastLine)))
	return res
}

func appendLine(words []string, currLen, maxWidth int) string {
	spaces := maxWidth - currLen

	if len(words) == 1 {
		surfix := strings.Repeat(" ", spaces)
		return words[0] + surfix
	}

	preCount := spaces % (len(words) - 1)  // 3%2=1
	avgSpaces := spaces / (len(words) - 1) // 3/2=1

	return strings.Join(words[:preCount+1], strings.Repeat(" ", avgSpaces+1)) + strings.Repeat(" ", avgSpaces) +
		strings.Join(words[preCount+1:], strings.Repeat(" ", avgSpaces))
}

// @lc code=end

