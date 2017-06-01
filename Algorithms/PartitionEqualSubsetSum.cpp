//
//  PartitionEqualSubsetSum.cpp
//  LeetCode
//  
//  URL: https://leetcode.com/problems/partition-equal-subset-sum
//  
//  Created by gongshang on 17/06/01.
//  Copyright (c) 2017年 ninerec. All rights reserved.
//

class Solution {
public:
    bool canPartition(vector<int>& nums) {
        sort(nums.begin(), nums.end(), less_equal<int>());
        int sum = accumulate(nums.begin(), nums.end(), 0);

        return !(sum % 2) && canPartition(nums, sum / 2 - nums[0], 1);
    }

    bool canPartition(vector<int> &nums, int sum, int start) {
        if (sum == 0) return true;
        if (start >= nums.size() || sum < nums[start]) return false;

        return  canPartition(nums, sum - nums[start], start + 1)
            || canPartition(nums, sum, start + 1);
    }
};
