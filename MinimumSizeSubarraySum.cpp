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
        int index = 0;
        int sum = 0;
        int length = 0;
        int ret = n;

        while (sum < s && index < n){
            sum += nums[index++];
            length++;
        }
        if (sum < s)
        {
            return 0;
        }

        ret = length;
        for (int i = 1; i < n; ++i)
        {
            sum -= nums[i-1];
            length--;

            while (sum < s && index < n) {
                sum += nums[index++];
                length++;       
            }

            if (sum >=s && ret > length)
            {
                ret = length;
            }
        }

        return ret;
    }
};
