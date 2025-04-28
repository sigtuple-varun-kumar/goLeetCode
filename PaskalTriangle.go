package leetcode

func generate(numRows int) [][]int {
	if numRows == 0 {
		return make([][]int, 0)
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	prevRows := generate(numRows - 1)

	prevRow := prevRows[len(prevRows)-1]

	currentRow := []int{1}

	for i := 1; i < numRows-1; i++ {
		currentRow = append(currentRow, prevRow[i-1]+prevRow[i])
	}
	currentRow = append(currentRow, 1)
	prevRows = append(prevRows, currentRow)
	return prevRows
}
