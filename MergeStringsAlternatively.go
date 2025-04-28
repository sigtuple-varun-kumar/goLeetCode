package leetcode

import "fmt"

func mergeAlternately(word1 string, word2 string) string {

	var ans string
	var i int
	for i = 0; i < len(word1) && i < len(word2); i++ {
		ans = ans + string(word1[i]) + string(word2[i])
	}
	if i < len(word1) {
		ans += word1[i:]
	} else if i < len(word2) {
		ans += word2[i:]
	}
	return ans
}

func testMergeAlternately() {
	fmt.Println(mergeAlternately("abc", "pqr"))
}
