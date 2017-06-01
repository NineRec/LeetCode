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
        int sum = accumulate(nums.begin(), nums.end(), 0);
        int target = sum / 2;
        vector<bool> dp(target + 1, 0);
        dp[0] = true;

        for (auto num : nums)
        {
            for (int i = target; i >= num; --i)
            {
                dp[i] = dp[i] || dp[i - num];
            }
        }

        return !(sum % 2) && dp[target];
    }
};
