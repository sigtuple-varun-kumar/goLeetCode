package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool { return i < j })
	prev := intervals[0]
	for i := 0; i < len(intervals); i++ {
		if prev[1] > intervals[i][0] {
			prev[1] = max(prev[1], intervals[i][1])
		} else {
			ans = append(ans, prev)
			prev = intervals[i]
		}
	}

	return append(ans, prev)
}
