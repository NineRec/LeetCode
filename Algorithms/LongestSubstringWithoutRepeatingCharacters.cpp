//
//  LongestSubstringWithoutRepeatingCharacters.cpp
//  LeetCode
//
//  Created by gongshang on 15/8/29.
//  Copyright (c) 2015年 ninerec. All rights reserved.
//

class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int index1 = 0, index2 = -1;
        int max = 0;
        unordered_map<char, int> hash;

        while (++index2 < s.length()) {
            auto temp = hash.find(s[index2]);
            if (temp != hash.end()) {
                index1 = (index1 > temp->second + 1) ? index1 : temp->second + 1;
            }
            max = (index2 - index1 >= max) ? index2 - index1 + 1 : max;
            hash[s[index2]] = index2; 
        }

        return max;
    }
};