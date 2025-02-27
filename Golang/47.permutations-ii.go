/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	slices.Sort(nums)
	// option 1: using a hash map to save existing result.
	// option 2: ignore the same value in the loop, use a hash map to record used
	return sortedPermuteUnique(nums)
}

func sortedPermuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	result := [][]int{}

	idx := 0
	for idx < len(nums) {
		idxVal := nums[idx]
		newNums := append(nums[:idx], nums[idx+1:]...)

		tmpRes := sortedPermuteUnique(newNums)
		for _, slc := range tmpRes {
			result = append(result, append([]int{idxVal}, slc...))
		}

		nums = append(newNums[:idx], append([]int{idxVal}, newNums[idx:]...)...)

		idx++
		for idx < len(nums) && nums[idx] == nums[idx-1] {
			idx++
		}
	}
	return result
}

// @lc code=end

