package golang

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	val, idx1, idx2 := 0, -1, -1
	for i := 0; i < (len(nums1)+len(nums2)+1)/2-1; i++ {
		_, idx1, idx2 = next(nums1, nums2, idx1, idx2)
	}

	val, idx1, idx2 = next(nums1, nums2, idx1, idx2)
	median := float64(val)
	if (len(nums1)+len(nums2))%2 == 0 {
		val, _, _ = next(nums1, nums2, idx1, idx2)
		median = (median + float64(val)) / 2
	}
	return median

}

func next(nums1 []int, nums2 []int, idx1, idx2 int) (int, int, int) {
	if idx1 == len(nums1)-1 {
		return nums2[idx2+1], idx1, idx2 + 1
	}
	if idx2 == len(nums2)-1 {
		return nums1[idx1+1], idx1 + 1, idx2
	}

	if nums1[idx1+1] < nums2[idx2+1] {
		return nums1[idx1+1], idx1 + 1, idx2
	} else {
		return nums2[idx2+1], idx1, idx2 + 1
	}
}

// @lc code=end
