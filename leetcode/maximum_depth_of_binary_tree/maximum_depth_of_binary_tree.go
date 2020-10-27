package maximum_depth_of_binary_tree

import "math"

// description
// https://leetcode.com/problems/maximum-depth-of-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	return int(recDepthSearch(root, 0.0))
}

func recDepthSearch(node *TreeNode, depth float64) float64 {
	if node == nil {
		return depth
	}
	depth++
	return math.Max(recDepthSearch(node.Left, depth), recDepthSearch(node.Right, depth))
}
