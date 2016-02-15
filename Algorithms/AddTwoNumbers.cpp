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
        if (l1 == NULL || l2 == NULL) {
            return NULL;
        }

        ListNode* head = new ListNode(-1);
        ListNode* ptr = head;
        int temp = 0;

        while (l1 != NULL && l2 != NULL) {
            int res = (l1->val + l2->val + temp) % 10;
            temp = (l1->val + l2->val + temp) / 10;
            ptr->next = new ListNode(res);
            ptr = ptr->next;

            l1 = l1->next;
            l2 = l2->next;
        }

        ListNode* rest = (l1 != NULL) ? l1 : l2;
        while (rest != NULL) {
            int res = (rest->val + temp) % 10;
            temp = (rest->val + temp) / 10;
            ptr->next = new ListNode(res);
            ptr = ptr->next;

            rest = rest->next;
        }

        if (temp != 0) {
            ptr->next = new ListNode(temp);
        }

        return head->next;
    }
};
