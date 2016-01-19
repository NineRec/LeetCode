//
//  BinaryTreeLevelOrderTraversal.cpp
//  LeetCode
//
//  Created by gongshang on 15/9/6.
//  Copyright (c) 2015年 ninerec. All rights reserved.
//

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> order;
        if (!root) return order;

        queue<TreeNode*> tree;
        tree.push(root);
        while(!tree.empty()) {
            int num = tree.size();
            vector<int> line;
            for (int i = 0; i < num; ++i) {
                TreeNode* node = tree.front();
                line.push_back(node->val);
                tree.pop();

                if (node->left){ tree.push(node->left); }
                if (node->right){ tree.push(node->right); }
            }
            order.push_back(line);
        }
        return order;
    }
};