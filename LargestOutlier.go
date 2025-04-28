package leetcode

import "math"

func GetLargestOutlier(nums []int) int {
	numSum := 0
	numMap := make(map[int]int)
	for _, num := range nums {
		numSum += num
		numMap[num]++
	}
	currOutlier := math.MinInt
	for _, num := range nums {
		potentialOutlier := numSum - 2*num
		if count := numMap[potentialOutlier]; count > 0 {
			if num == potentialOutlier {
				if count > 1 {
					currOutlier = max(potentialOutlier, currOutlier)
				}
			} else {
				currOutlier = max(potentialOutlier, currOutlier)
			}
		}
	}
	return currOutlier
}
