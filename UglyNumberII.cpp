//
//  UglyNumber.cpp
//  LeetCode
//
//  Created by gongshang on 15/9/4.
//  Copyright (c) 2015年 ninerec. All rights reserved.
//

class Solution {
public:
    int nthUglyNumber(int n) {
        if (n == 1) {
            return 1;
        }

        int ugly[n];
        ugly[0] = 1;

        int index2 = 0;
        int index3 = 0;
        int index5 = 0;

        for (int i = 1; i < n; ++i){
            if ((ugly[index2] * 2 < ugly[index3] * 3) && (ugly[index2] * 2 < ugly[index5] * 5)){
                ugly[i] = ugly[index2] * 2;
            } else {
                ugly[i] = (ugly[index3] * 3 < ugly[index5] * 5)? ugly[index3] * 3: ugly[index5] * 5;
            }

            if (ugly[index2] * 2 == ugly[i]){
                index2++;
            }

            if (ugly[index3] * 3 == ugly[i]){
                index3++;
            }

            if (ugly[index5] * 5 == ugly[i]){
                index5++;
            }
        }

        return ugly[n-1];
    }
};
