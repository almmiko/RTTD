package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root != nil {
		maxLeft := maxDepth(root.Left) + 1
		maxRight := maxDepth(root.Right) + 1

		if maxLeft > maxRight {
			return maxLeft
		} else {
			return maxRight
		}

	}

	return 0
}
