// 665. Non-decreasing Array
// Given an array with n integers, your task is to check if it could become non-decreasing by
// modifying at most 1 element.
//
// We define an array is non-decreasing if array[i] <= array[i + 1] holds for every i (1 <= i < n).

package golang

func checkPossibility(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return isNonDecreasing(removeIndex(nums, i)) || isNonDecreasing(removeIndex(nums, i+1))
		}
	}
	return true
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0, len(s))
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func isNonDecreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
