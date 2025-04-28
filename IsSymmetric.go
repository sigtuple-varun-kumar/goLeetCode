package leetcode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return walk(root.Left, root.Right)
}

func walk(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return walk(root1.Left, root2.Right) && walk(root1.Right, root2.Left)
}
