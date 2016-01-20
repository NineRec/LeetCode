//
//  HIndex.cpp
//  LeetCode
//
//  Created by gongshang on 15/9/6.
//  Copyright (c) 2015年 ninerec. All rights reserved.
//

class Solution {
public:
    int hIndex(vector<int>& citations) {
        int size = citations.size();

        int count[size + 1];
        memset(count, 0, (size + 1)*sizeof(int));

        for (std::vector<int>::iterator i = citations.begin(); i != citations.end(); ++i){
            int index = (*i > size)?size: *i;
            count[index] += 1;
        }

        int hindex = 0;
        for (int i = size; i >= 0; --i){
            hindex = hindex + count[i];
            if (hindex >= i){
                return i;
            }
        }
        return 0;
    }
};