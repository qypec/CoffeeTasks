package binary_tree_level_order_traversal

// description
// https://leetcode.com/problems/binary-tree-level-order-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := make([][]int, 0)
	queue := make([]*TreeNode, 0, 1)
	queue = append(queue, root)
	for len(queue) != 0 {
		levelNodes := make([]int, 0, len(queue)*2)
		queueCopy := queue[:]
		for i := 0; i < len(queueCopy); i++ {
			if queueCopy[i].Left != nil {
				queue = append(queue, queueCopy[i].Left)
			}
			if queueCopy[i].Right != nil {
				queue = append(queue, queueCopy[i].Right)
			}
			levelNodes = append(levelNodes, queueCopy[i].Val)
		}
		queue = queue[len(queueCopy):]
		res = append(res, levelNodes)
	}
	return res
}
