//
//  TwoSum.cpp
//  LeetCode
//
//  Created by gongshang on 16/02/16.
//  Copyright (c) 2015年 ninerec. All rights reserved.
//

class Solution {
public:
    string convert(string s, int numRows) {
        if (numRows == 1 || s.length() <= numRows) return s;
        
        string result(s);
        int result_index = 0;

        for (int row = 0; row < numRows && row < s.length(); ++row) {
            int space0 = (numRows - row) * 2 - 2;
            int space1 = row * 2;
            
            int pos = row;
            result[result_index++] = s[pos];
            while (pos < s.length()) {
                pos = pos + space0;
                if (space0 && pos < s.length()) {
                    result[result_index++] = s[pos];
                }
                pos = pos + space1;
                if (space1 && pos < s.length()) {
                    result[result_index++] = s[pos];
                }
            }
        }

        return result;
    }
};