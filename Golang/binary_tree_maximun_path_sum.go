// Given a non-empty binary tree, find the maximum path sum.
// For this problem, a path is defined as any sequence of nodes from some starting node to any node
// in the tree along the parent-child connections. The path must contain at least one node and
// does not need to go through the root.

package golang

var maxSum int

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	maxSum = root.Val
	maxChildSum(root)
	return maxSum
}

func maxChildSum(root *TreeNode) int {
	// the max sum of nil
	if root == nil {
		return 0
	}

	valLeft := max(maxChildSum(root.Left), 0)
	valRight := max(maxChildSum(root.Right), 0)

	maxSum = max(maxSum, root.Val+valLeft+valRight)
	return root.Val + max(valLeft, valRight)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
