package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	solution := [][]int{}
	n := len(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip duplicates
		}

		left, right := i+1, n-1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				solution = append(solution, []int{nums[i], nums[left], nums[right]})
				// Skip duplicates
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return solution
}
