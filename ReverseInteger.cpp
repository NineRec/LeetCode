//
//  ReverseInteger.cpp
//  LeetCode
//
//  Created by gongshang on 15/9/2.
//  Copyright (c) 2015年 ninerec. All rights reserved.
//

class Solution {
public:
    int reverse(int x) {
        int sign;
        sign = (x<0)? -1: 1;

        x = abs(x);
        int result = 0;
        int bit = 0;    // 个位
        while (x > 0) {
            bit = x % 10;
            x = x / 10;
            if (x == 0 && (INT_MAX - bit +0.0) / 10 < result) {
                return 0;
            }
            
            result = result * 10 + bit;
        }

        return sign * result;
    }
};
