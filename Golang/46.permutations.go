package golang

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	resLen := 1
	for i := 1; i <= len(nums); i++ {
		resLen = resLen * i
	}

	output := make([][]int, 0, resLen)

	for i, num := range nums {
		nums[0], nums[i] = nums[i], nums[0] // swap to index 0
		subOutput := permute(nums[1:])
		nums[0], nums[i] = nums[i], nums[0] // swap back

		for _, ints := range subOutput {
			output = append(output, append([]int{num}, ints...))
		}
	}
	return output
}

// @lc code=end
