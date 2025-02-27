package golang

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// odd  => idx: (len(nums1)+len(nums2))/2 +1
	// even => idx: (len(nums1)+len(nums2))/2, (len(nums1)+len(nums2))/2+1
	// m/2 n/2 <= k

	// binary search =>
	if (len(nums1)+len(nums2))%2 == 1 {
		return float64(search(nums1, nums2, (len(nums1)+len(nums2))/2+1))
	}

	kthV := float64(search(nums1, nums2, (len(nums1)+len(nums2))/2))
	nextV := float64(search(nums1, nums2, (len(nums1)+len(nums2))/2+1))
	return (kthV + nextV) / 2
}

func search(nums1 []int, nums2 []int, kth int) int {
	if len(nums1) == 0 {
		return nums2[kth-1]
	}
	if len(nums2) == 0 {
		return nums1[kth-1]
	}

	if kth == 1 {
		return min(nums1[0], nums2[0])
	}

	mid1 := min(kth/2-1, len(nums1)-1)
	mid2 := min(kth/2-1, len(nums2)-1)
	if nums1[mid1] < nums2[mid2] {
		return search(nums1[mid1+1:], nums2, kth-mid1-1)
	} else {
		return search(nums1, nums2[mid2+1:], kth-mid2-1)
	}
}

// @lc code=end
