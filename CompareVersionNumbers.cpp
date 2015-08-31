//
//  versionCompare.cpp
//  LeetCode
//
//  Created by gongshang on 15/8/31.
//  Copyright (c) 2015年 ninerec. All rights reserved.
//

class Solution {
public:
    vector<int> &split(const string &s, char delim, vector<int> &elems) {
        stringstream ss(s);
        string item;
        int num;
        while (getline(ss, item, delim)) {
            elems.push_back(atoi(item.c_str()));
        }
        return elems;
    }

    vector<int> split(const string &s, char delim) {
        vector<int> elems;
        split(s, delim, elems);
        return elems;
    }

    void trim_zero(vector<int>& v){
        while (v.size() > 1 && v.back() == 0){
            v.pop_back();
        }
    }

    int compareVersion(string version1, string version2) {
        int ret = 0;
        vector<int> nums1 = split(version1, '.');
        vector<int> nums2 = split(version2, '.');

        trim_zero(nums1);
        trim_zero(nums2);

        for (int i = 0; i < nums1.size() && i < nums2.size(); ++i)
        {
            if (nums1[i] != nums2[i])
            {
                return (nums1[i] > nums2[i])? 1 : -1;
            }
        }

        if (nums1.size() == nums2.size())
        {
            return 0;
        }
        return (nums1.size() > nums2.size())? 1 : -1;
    }
};
