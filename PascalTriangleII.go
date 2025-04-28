package leetcode

func getRow(rowIndex int) []int {
	triangle := getPrevRows(rowIndex)
	return triangle[rowIndex]
}

func getPrevRows(numRows int) [][]int {
	if numRows == 0 {
		return make([][]int, 0)
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	prevRows := getPrevRows(numRows - 1)

	prevRow := prevRows[len(prevRows)-1]

	currRow := []int{1}

	for i := 1; i < numRows-1; i++ {
		currRow = append(currRow, prevRow[i-1]+prevRow[i])
	}
	currRow = append(currRow, 1)
	prevRows = append(prevRows, currRow)
	return prevRows
}
