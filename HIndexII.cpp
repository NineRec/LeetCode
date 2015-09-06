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
        int index = 0;

        int size = citations.size();
        while (++index <= size && citations[size - index] >= index) {}
        
        return index;
    }
};