package leetcode

func trap(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	totalWater := 0
	left, right := 0, len(heights)-1
	maxLeft, maxRight := 0, 0
	for left < right {
		if heights[left] < heights[right] {
			if heights[left] >= maxLeft {
				maxLeft = heights[left]
			} else {
				totalWater += maxLeft - heights[left]
			}
			left++
		} else {
			if heights[right] > maxRight {
				maxRight = heights[right]
			} else {
				totalWater += maxRight - heights[right]
			}
			right--
		}
	}
	return totalWater
}
