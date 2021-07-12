package golang

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	leftInorder, rightInorder := seprateByRootVal(inorder, preorder[0])
	leftPreorder := preorder[1 : len(leftInorder)+1]
	rightPreorder := preorder[len(leftInorder)+1:]

	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)

	return root
}

func seprateByRootVal(inorder []int, rootVal int) (leftInorder []int, rightInorder []int) {
	for idx, v := range inorder {
		if v == rootVal {
			return inorder[:idx], inorder[idx+1:]
		}
	}

	return nil, nil
}

// @lc code=end
