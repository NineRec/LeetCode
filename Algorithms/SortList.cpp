//
//  SortList.cpp
//  LeetCode
//
//  Created by gongshang on 15/10/22.
//  Copyright (c) 2015年 ninerec. All rights reserved.
//

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* sortList(ListNode* head) {
        int length = count_size(head);
        
        ListNode virtual_head = new ListNode(0);
        virtual_head->next = head;
        
        int block = 1;
        while (block < length) {
            
            

            block *= 2;
        }
        
        return virtual_head->next;
    }

private:
    int count_size(ListNode* head) {
        int total = 0;
        while (head != NULL) {
            total++;
            head = head->next;
        }

        return total;
    }
};