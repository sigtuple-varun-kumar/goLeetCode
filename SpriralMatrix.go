package leetcode

func spiralOrder(matrix [][]int) []int {
	top, bottom := 0, len(matrix)
	left, right := 0, len(matrix[0])
	ansList := make([]int, 0)

	for top <= bottom && left <= right {
		// traverse right
		for i := left; i < right; i++ {
			ansList = append(ansList, matrix[top][i])
		}
		//traverse down increment top by 1
		top++
		for i := top; i < bottom; i++ {
			ansList = append(ansList, matrix[i][right])
		}
		// traverse from right to left decrement right by 1
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				ansList = append(ansList, matrix[bottom][i])
			}
		}

		// traverse from bottom to top and decrement bottom
		bottom--
		if left <= right {
			for i := bottom; i >= top; i-- {
				ansList = append(ansList, matrix[left][i])
			}
		}
		left++
	}
	return ansList
}
