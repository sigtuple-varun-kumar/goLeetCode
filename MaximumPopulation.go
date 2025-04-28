package leetcode

import "sort"

func maximumPopulation(logs [][]int) int {
	freq := make(map[int]int)
	for _, log := range logs {
		freq[log[0]] += 1
		freq[log[1]] -= 1
	}
	years := make([]int, 0)
	for year := range freq {
		years = append(years, year)
	}
	sort.Ints(years)
	max_p := 0
	ans_year := 0
	count := 0

	for _, year := range years {
		count += freq[year]
		if count > max_p {
			max_p = count
			ans_year = year
		}
	}
	return ans_year
}
