package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
}

func depthFirstSearch(node *TreeNode) int {
	if node.Left == nil && node.Right == nil {
		return 1
	}
	var leftDepth int
	var rightDepth int
	if node.Left != nil {
		leftDepth = 1 + depthFirstSearch(node.Left)
	}
	if node.Right != nil {
		rightDepth = 1 + depthFirstSearch(node.Right)
	}
	return int(math.Max(float64(leftDepth), float64(rightDepth)))
}
