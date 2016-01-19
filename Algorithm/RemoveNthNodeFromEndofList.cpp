//
//  MinimumWindowSubstring.cpp
//  LeetCode
//
//  Created by gongshang on 15/9/2.
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
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode* ptr = head;
        ListNode* del_prt = head;

        int i =0;
        for (; i < n && ptr->next != NULL; ++i){
            ptr = ptr->next;    
        }

        while (ptr->next != NULL){
            ptr = ptr->next;
            del_prt = del_prt->next;
        }

        if (i < n && del_prt == head)    // delete head
        {
            head = head->next;
            free(del_prt);
        } else {
            ptr = del_prt->next;
            del_prt->next = ptr->next;
            free(ptr);
        }
        return head;
    }
};