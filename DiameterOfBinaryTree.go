package leetcode

func diameterOfBinaryTree(root *TreeNode) int {
	maximumDiameter := 0
	findDiameter(root, &maximumDiameter)
	return maximumDiameter
}

func findDiameter(node *TreeNode, maximumDiameter *int) int {
	if node == nil {
		return 0
	}
	LeftDiameter := findDiameter(node.Left, maximumDiameter)
	RightDiameter := findDiameter(node.Left.Right, maximumDiameter)
	*maximumDiameter = max(*maximumDiameter, LeftDiameter+RightDiameter)
	return 1 + max(LeftDiameter, RightDiameter)
}
