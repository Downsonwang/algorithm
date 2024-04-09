/*
 * @Descripttion: 两个不相交的数据结构
 * @Author:DW
 * @Date: 2024-04-07 21:45:15
 * @LastEditTime: 2024-04-08 22:18:28
 */
package unionfind

type UnionFind struct {
	parent []int
	// 秩的作用就是通过记录树的高度（或者近似的高度），在合并操作时优化树的高度，从而尽可能保持树的平衡
	// ，避免树的高度过高。通过将较短的树链接到较高的树上，可以减少整个树的高度，
	// 从而提高查找操作（find）和合并操作（union）的效率。
	rank []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
	}

	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 0
	}

	return uf
}

// 保持并查集平衡 防止树过高
func (uf *UnionFind) union(x, y int) {
	rx, ry := uf.find(x), uf.find(y)
	if rx != ry {
		if uf.rank[rx] < uf.rank[ry] {
			uf.parent[rx] = ry
		} else if uf.rank[rx] > uf.rank[ry] {
			uf.parent[ry] = rx
		} else {
			uf.parent[ry] = rx
			uf.rank[rx]++
		}
	}
}

// 路径压缩  +  减少树高 拽出
func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {

		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func FindConnectedComponents(graph map[int][]int) [][]int {
	numVer := len(graph)
	uf := NewUnionFind(numVer)
	for u := range graph {
		for _, v := range graph[u] {
			if uf.find(u) != uf.find(v) {
				uf.union(u, v)
			}
		}
	}

	components := make(map[int][]int)

	for i := 0; i < numVer; i++ {
		root := uf.find(i)
		components[root] = append(components[root], i)
	}

	var result [][]int
	for _, vertices := range components {
		result = append(result, vertices)
	}
	return result

}
