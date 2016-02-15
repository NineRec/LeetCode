//
//  TwoSum.cpp
//  LeetCode
//
//  Created by gongshang on 16/02/14.
//  Copyright (c) 2015年 ninerec. All rights reserved.
//

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> hash;

        for (int index = 0; index < nums.size(); ++index)
        {
            auto result = hash.find(target - nums[index]);
            if (result != hash.end())
            {
                return {hash[target - nums[index]], index};
            }
            hash[nums[index]] = index;
        }
    }
};
