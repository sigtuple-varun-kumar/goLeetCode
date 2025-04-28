package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	walkForInvert(root)
	return root
}

func walkForInvert(root *TreeNode) {
	if root == nil {
		return
	}
	walkForInvert(root.Left)
	walkForInvert(root.Right)
	root.Left, root.Right = root.Right, root.Left
}
