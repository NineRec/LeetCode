//
//  BinarySearchTree.cpp
//  LeetCode
//
//  Created by gongshang on 15/8/25.
//  Copyright (c) 2015年 ninerec. All rights reserved.
//

/**
 * Definition for binary tree
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

class BSTIterator {
public:
    stack<TreeNode*> *stk;
    BSTIterator(TreeNode *root) {
        stk = new stack<TreeNode*>();
        while (root) {
            stk->push(root);
            root = root->left;
        }
    }
    
    /** @return whether we have a next smallest number */
    bool hasNext() {
        return !stk->empty();
    }
    
    /** @return the next smallest number */
    int next() {
        TreeNode* node = stk->top();
        stk->pop();
        
        int ret = node->val;
        
        if (node->right) {
            node = node->right;
            while (node) {
                stk->push(node);
                node = node->left;
            }
        }
        return ret;
    }
};

/**
 * Your BSTIterator will be called like this:
 * BSTIterator i = BSTIterator(root);
 * while (i.hasNext()) cout << i.next();
 */