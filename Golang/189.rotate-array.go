/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	prev := make([]int, k)
	copy(prev, nums[len(nums)-k:])
	for i := len(nums) - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}
	for i, v := range prev {
		nums[i] = v
	}
}

// @lc code=end

