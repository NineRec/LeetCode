//
//  AddTwoNumbers.cpp
//  LeetCode
//
//  Created by gongshang on 15/8/29.
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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode head = ListNode(-1);
        ListNode* ptr = &head;

        int temp = 0;
        while (l1 || l2 || temp) {
            int sum = (l1 ? l1->val : 0) + (l2 ? l2->val : 0) + temp;
            temp = sum / 10;
            ptr->next = new ListNode(sum % 10);
            ptr = ptr->next;

            l1 = l1 ? l1->next : l1;
            l2 = l2 ? l2->next : l2;
        }

        return head.next;
    }
};
