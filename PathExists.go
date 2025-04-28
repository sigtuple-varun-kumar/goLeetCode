package leetcode

func validPath(n int, edges [][]int, source int, destination int) bool {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}

	for _, edge := range edges {
		union(edge, rank, parent)
	}
	return find(source, parent) == find(destination, parent)
}

func find(x int, parent []int) int {
	if parent[x] != x {
		parent[x] = find(parent[x], parent)
	}
	return parent[x]
}

func union(edge []int, rank []int, parent []int) {
	rootX := find(edge[0], parent)
	rootY := find(edge[1], parent)
	if rootX != rootY {
		if rank[rootX] > rank[rootY] {
			parent[rootY] = rootX
		} else if rank[rootY] > rank[rootX] {
			parent[rootX] = rootY
		} else {
			parent[rootY] = rootX
			rank[rootX]++
		}
	}
}
