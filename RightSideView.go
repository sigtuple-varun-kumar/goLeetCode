package leetcode

func rightSideView(root *TreeNode) []int {

	// bfs approach
	// if root == nil {
	// 	return nil
	// }
	// queue := []*TreeNode{root}
	// ans := make([]int, 0)

	// for len(queue) > 0 {

	// 	childNodes := make([]*TreeNode, 0)
	// 	var prev *TreeNode
	// 	for len(queue) > 0 {
	// 		currentNode := queue[0]
	// 		queue = queue[1:]
	// 		if currentNode.Left != nil {
	// 			childNodes = append(childNodes, currentNode.Left)
	// 		}
	// 		if currentNode.Right != nil {
	// 			childNodes = append(childNodes, currentNode.Right)
	// 		}
	// 		prev = currentNode
	// 	}
	// 	ans = append(ans, prev.Val)
	// 	queue = childNodes[:]
	// }
	// return ans

	// dfs managing depth is much easier with depth first search or recurson

	ans := make([]int, 0)
	rightDFS(root, &ans, 0)
	return ans
}

func rightDFS(node *TreeNode, ans *[]int, depth int) {
	if node == nil {
		return
	}
	if depth == len(*ans) {
		*ans = append(*ans, node.Val)
	}
	rightDFS(node.Right, ans, depth+1)
	rightDFS(node.Left, ans, depth+1)
}
