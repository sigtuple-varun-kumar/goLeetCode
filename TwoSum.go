package leetcode

func twoSum(nums []int, target int) []int {

	var numMap = make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if _, ok := numMap[complement]; ok {
			return []int{i, numMap[complement]}
		}
		numMap[num] = i
	}
	return []int{}
}
