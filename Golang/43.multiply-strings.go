package golang

import "strconv"

/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	digits1 := make([]int, len(num1))
	for i, v := range num1 {
		digit, _ := strconv.Atoi(string(v))
		digits1[i] = digit
	}

	res := make([]int, len(num1)+len(num2))
	flag, k := 0, 0
	for i := len(num2) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(num2[i]))

		for j := len(digits1) - 1; j >= 0; j-- {
			tmp := digit*digits1[j] + flag + res[k]

			res[k] = tmp % 10
			flag = tmp / 10

			k++
		}

		for flag != 0 {
			tmp := flag + res[k]

			res[k] = tmp % 10
			flag = tmp / 10

			k++
		}

		k = len(num2) - i
	}

	return convert(res)
}

func convert(nums []int) string {
	idx := len(nums) - 1
	for ; nums[idx] == 0; idx-- {
		continue
	}

	res := ""
	for ; idx >= 0; idx-- {
		res += strconv.FormatInt(int64(nums[idx]), 10)
	}

	return res
}

// @lc code=end
