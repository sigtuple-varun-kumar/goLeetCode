package leetcode

import (
	"fmt"
	"sort"
)

func NoOfPointsI(nums [][]int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})
	mergedSlice := make([][]int, 0)
	prev := nums[0]
	fmt.Println(nums)
	for _, num := range nums[1:] {
		if num[0] <= prev[1] {
			prev[1] = max(num[1], prev[1])
		} else {
			mergedSlice = append(mergedSlice, prev)
			prev = num
		}
	}
	mergedSlice = append(mergedSlice, prev)
	fmt.Println(mergedSlice)
	numberOfPoints := 0
	for _, interval := range mergedSlice {
		numberOfPoints += interval[1] - interval[0] + 1
	}
	return numberOfPoints
}

func NoOfPointsII(nums [][]int) int {
	var cars [102]int
	for _, num := range nums {
		cars[num[0]]++
		cars[num[1]+1]--
	}
	totalPoints, noOfCars := 0, 0

	for _, car := range cars {
		noOfCars += car
		if noOfCars > 0 {
			totalPoints++
		}
	}
	return totalPoints
}
