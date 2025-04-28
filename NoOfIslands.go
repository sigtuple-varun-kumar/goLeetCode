package leetcode

func numIslands(grid [][]byte) int {
	number := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				number++
				dfs(&grid, i, j)
			}
		}
	}
	return number
}

func dfs(grid *[][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) {
		return
	}
	// no need to do anything if we have 0
	if (*grid)[i][j] == '0' {
		return
	}
	// alll the recursion base cases are done, now need to mark node as visited
	(*grid)[i][j] = '0'

	// and recurse or do dfs

	dfs(grid, i-1, j) // top
	dfs(grid, i+1, j) // bottom
	dfs(grid, i, j-1) // left
	dfs(grid, i, j+1) // right
}
