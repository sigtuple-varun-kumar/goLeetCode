package leetcode

func isZeroArray(nums []int, queries [][]int) bool {
	freq := make([]int, len(nums)+1)

	for _, query := range queries {
		freq[query[0]]++
		freq[query[1]+1]--
	}
	currFreq := 0
	for i, num := range nums {
		currFreq += freq[i]
		if currFreq < num {
			return false
		}
	}
	return true
}
