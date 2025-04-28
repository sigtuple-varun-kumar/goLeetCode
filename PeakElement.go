package leetcode

import (
	"fmt"
	"math"
)

func FindPeakElement(nums []int) int {
	return divideAndConquer(nums, 0, len(nums)-1)
}

func divideAndConquer(nums []int, start int, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	var leftElement int
	if mid-1 < 0 {
		leftElement = math.MinInt
	} else {
		leftElement = nums[mid-1]
	}

	var rightElement int
	if mid+1 >= len(nums) {
		rightElement = math.MinInt
	} else {
		rightElement = nums[mid+1]
	}

	if nums[mid] > leftElement && nums[mid] > rightElement {
		fmt.Println("Found it", mid)
		fmt.Println("left element is ", leftElement)
		fmt.Println("right element is ", rightElement)
		return mid
	}

	leftPeak := divideAndConquer(nums, start, mid-1)
	if leftPeak > 0 && leftPeak < len(nums) {
		return leftPeak
	}

	rightPeak := divideAndConquer(nums, mid+1, end)
	if rightPeak > 0 && rightPeak < len(nums) {
		return rightPeak
	}
	return -1
}
