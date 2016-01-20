//
//  MinimumSizeSubarraySum.cpp
//  LeetCode
//
//  Created by gongshang on 15/8/29.
//  Copyright (c) 2015年 ninerec. All rights reserved.
//

class Solution {
public:
    int minSubArrayLen(int s, vector<int>& nums) {
        int n = nums.size();
        int min = n+1;
        int begin = 0;
        int sum = 0;

        for (int i = 0; i < n; ++i)
        {
            sum += nums[i];

            while (sum >= s && begin <= i) {
                if ((i - begin + 1) < min)
                {
                    min = i - begin + 1;
                }
                sum -= nums[begin];
                begin++;
            }
        }

        return (min>n)? 0: min;
    }
};
