//
//  DivideTwoIntegers.cpp
//  LeetCode
//
//  Created by gongshang on 15/9/6.
//  Copyright (c) 2015年 ninerec. All rights reserved.
//

class Solution {
public:
    int divide(int dividend, int divisor) {
        if (divisor == 0) return INT_MAX;
        if (dividend == INT_MIN && divisor== -1) return INT_MAX;

        int sign = (dividend >> 31) ^ (divisor >> 31);
        long dvd = labs(dividend);
        long dvs = labs(divisor);

        int ans = 0;
        while (dvd >= dvs){
            int temp = 0;
            while (dvd >= (dvs << temp + 1)){
                temp ++;
            }
            dvd -= dvs << temp;
            ans += 1 << temp;
        }

        return (sign==0)? ans: -ans;
    }
};
