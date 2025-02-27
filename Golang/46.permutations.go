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

	output := [][]int{}
	for i := 0; i < len(nums); i++ {
		nums[0], nums[i] = nums[i], nums[0]

		subRes := permute(nums[1:])
		for _, slc := range subRes {
			output = append(output, append(slc, nums[0]))
		}

		nums[0], nums[i] = nums[i], nums[0]
	}
	return output
}

// @lc code=end
