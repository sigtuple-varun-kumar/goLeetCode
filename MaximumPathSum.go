/*
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package leetcode

import "math"

func maxPathSum(root *TreeNode) int {
	maximumSum := int(math.MinInt)
	findMaxSum(root, &maximumSum)
	return maximumSum
}

func findMaxSum(node *TreeNode, maximumSum *int) int {
	if node == nil {
		return 0
	}
	leftSum := findMaxSum(node.Left, maximumSum)
	rightSum := findMaxSum(node.Right, maximumSum)

	if leftSum < 0 {
		leftSum = 0
	}
	if rightSum < 0 {
		rightSum = 0
	}
	*maximumSum = max(*maximumSum, leftSum+node.Val+rightSum)

	return node.Val + max(leftSum, rightSum)
}
