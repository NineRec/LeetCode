//
//  MinimumWindowSubstring.cpp
//  LeetCode
//
//  Created by gongshang on 15/9/1.
//  Copyright (c) 2015年 ninerec. All rights reserved.
//

class Solution {
private:
    const int MAX_INTEGER = 2147483647;

    map<char, int> t_map;
    map<char, int> s_map;

public:
    string minWindow(string s, string t) {
        int min = MAX_INTEGER;
        int start;
        int winSize = t.length();
        int curSize = 0;

        for (int i = 0; i < t.length(); ++i){
            if (t_map.find(t[i]) != t_map.end())
            {
                t_map[t[i]] = t_map[t[i]] + 1;
            } else {
                t_map[t[i]] = 1;
                s_map[t[i]] = 0;
            }
        }

        int temp_start = -1;
        for (int i = 0; i < s.length(); ++i){
            if (t_map.find(s[i]) != t_map.end()){   // in the map
                s_map[s[i]] = s_map[s[i]] + 1;  // add s_map

                if (curSize < winSize && s_map[s[i]] <= t_map[s[i]])
                {
                    curSize++;
                }

                if (temp_start == -1)
                {
                    temp_start = i;
                }

                if (curSize >= winSize) 
                {
                    while(s_map[s[temp_start]] > t_map[s[temp_start]]){ //去掉重复
                        s_map[s[temp_start]]--;
                        while (t_map.find(s[++temp_start]) == t_map.end()){}
                    }
                    
                    if (min > (i - temp_start + 1))
                    {
                        start = temp_start;
                        min = i - temp_start + 1;
                    }            
                }
            }
        }

        return (curSize < winSize)?"":s.substr(start, min);
    }
};