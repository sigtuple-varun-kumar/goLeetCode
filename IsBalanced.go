package leetcode

import "math"

// public class Solution {
// 	private boolean result = true;

// 	public boolean isBalanced(TreeNode root) {
// 		maxDepth(root);
// 		return result;
// 	}

// 	public int maxDepth(TreeNode root) {
// 		if (root == null)
// 			return 0;
// 		int l = maxDepth(root.left);
// 		int r = maxDepth(root.right);
// 		if (Math.abs(l - r) > 1)
// 			result = false;
// 		return 1 + Math.max(l, r);
// 	}
// 	}

func isBalanced(root *TreeNode) bool {
	result := true
	var maxDepth func(node *TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := maxDepth(node.Left)
		r := maxDepth(node.Right)
		if math.Abs(float64(l-r)) > 1 {
			result = false
		}
		return 1 + int(math.Max(float64(l), float64(r)))
	}
	maxDepth(root)
	return result
}
