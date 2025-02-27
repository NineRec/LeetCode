package golang

/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	// Insert in word1 == Delete in word2 => only Insert & Replace.
	// f(i,j) = f(i-1,j-1)+0(==), f(i-1,j-1)+1, f(i,j-1)+1, f(i-1,j)+1
	// for range desc

	minDist := make([]int, len(word2)+1)
	for key := range minDist {
		minDist[key] = key
	}

	for i := 1; i <= len(word1); i++ {
		prev, prev1 := minDist[0], -1
		minDist[0] = i

		for j := 1; j <= len(word2); j++ {
			prev1, prev = prev, minDist[j] // i-1,j-1, i-1,j

			if word1[i-1] == word2[j-1] {
				minDist[j] = min(prev1, minDist[j-1]+1, prev+1)
			} else {
				minDist[j] = min(prev1+1, minDist[j-1]+1, prev+1)
			}
		}
	}

	return minDist[len(word2)]
}

// @lc code=end
