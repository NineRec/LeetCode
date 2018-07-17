package Golang

import (
    "fmt"
)

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := hash[target - nums[i]]; ok {
			return []int {j, i}
		}
		hash[nums[i]] = i
	}
	return []int {-1, len(nums)}
}

func main() {
    fmt.Print(twoSum([]int {3, 2, 4}, 6))
}
