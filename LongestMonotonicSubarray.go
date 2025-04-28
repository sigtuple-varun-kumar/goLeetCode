package leetcode

import "fmt"

func LongestMonotonicSubarray(nums []int) int {
	largestLength := 1
	currentIncreasingLen := 1
	currentDecreasingLen := 1

	for i := 1; i < len(nums); i++ {
		fmt.Println(nums[i-1], nums[i])
		if nums[i] == nums[i-1] {
			currentDecreasingLen = 1
			currentIncreasingLen = 1
		} else if nums[i-1] < nums[i] {
			currentIncreasingLen++
			currentDecreasingLen = 1
		} else {
			currentDecreasingLen++
			currentIncreasingLen = 1
		}
		largestLength = max(largestLength, currentDecreasingLen, currentIncreasingLen)
	}
	return largestLength
}
