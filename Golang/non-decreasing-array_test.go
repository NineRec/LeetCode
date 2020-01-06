// 665. Non-decreasing Array
// Given an array with n integers, your task is to check if it could become non-decreasing by
// modifying at most 1 element.
//
// We define an array is non-decreasing if array[i] <= array[i + 1] holds for every i (1 <= i < n).

package golang

import "testing"

func Test_checkPossibility(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "3,4,5,2",
			nums: []int{3, 4, 5, 2},
			want: true,
		},
		{
			name: "4,2,1",
			nums: []int{4, 2, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibility(tt.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
