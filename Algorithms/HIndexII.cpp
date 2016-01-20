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
        if (citations.empty()) return 0;

        int size = citations.size();
        int left = 0;
        int right = size - 1;
        int middle;

        while (left < right) {
            middle = (left + right) / 2;
            if (citations[middle] >= size - middle){
                right = middle;
            } else {
                left = middle + 1;
            }
        }
        return (citations[left]>=size-left)?size - left:size-left-1;
    }
};