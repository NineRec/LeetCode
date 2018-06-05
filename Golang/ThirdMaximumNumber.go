// Given a non-empty array of integers, return the third maximum number 
// in this array. If it does not exist, return the maximum number. 
// The time complexity must be in O(n).
package Golang

func thirdMax(nums []int) int {
	minInt := -(int(^uint(0) >> 1)) - 1

	top1 := nums[0]
	top2 := minInt
	top3 := minInt
	for i := 1; i < len(nums); i++ {
		if nums[i] == top1 || nums[i] == top2 || nums[i] <= top3 {
			continue
		}

		top3 = nums[i]

		if top3 > top2 {
			top3, top2 = top2, top3
		}
		if top2 > top1 {
			top2, top1 = top1, top2
		}
	}

	if top3 == minInt {
		return top1
	}
	return top3
}
