package leetcode

// sample example = [[1,2],[1,1],[3,4],[4,5],[5,6],[7,7]]©
// return number of connectred nodes
import "fmt"

func NumberOfComponents(properties [][]int, k int) int {
	noOfConnectedNodes := 0
	for i := 0; i < len(properties)-1; i++ {
		if intersect(properties[i], properties[i+1]) > k {
			noOfConnectedNodes++
		}
	}
	return noOfConnectedNodes
}

func intersect(a []int, b []int) int {
	commonElements := 0
	elementsMap := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		elementsMap[a[i]] = true
	}

	for i := 0; i < len(b); i++ {
		if elementsMap[b[i]] {
			commonElements++
		}
	}
	fmt.Println(commonElements)
	return commonElements
}
