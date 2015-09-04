//
//  UglyNumber.cpp
//  LeetCode
//
//  Created by gongshang on 15/9/4.
//  Copyright (c) 2015年 ninerec. All rights reserved.
//

class Solution {
public:
    bool isUgly(int num) {
        if (num <= 0)
        {
            return false;
        }

        while (num % 2 == 0) num = num / 2; 
        while (num % 3 == 0) num = num / 3;
        while (num % 5 == 0) num = num / 5;

        return (num == 1)? true: false;
    }
};
