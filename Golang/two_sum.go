// 1. Two Sum
// Given an array of integers, return indices of the two numbers such that they add up to
// a specific target.
// You may assume that each input would have exactly one solution, and you may not use the
// same element twice.

package golang

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := hash[target-nums[i]]; ok {
			return []int{j, i}
		}
		hash[nums[i]] = i
	}
	return []int{-1, len(nums)}
}
