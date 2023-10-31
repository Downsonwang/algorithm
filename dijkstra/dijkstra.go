/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-10-30 21:13:23
 * @LastEditTime: 2023-10-31 12:39:24
 */
package dijkstra

import (
	"math"
)

// djkstra - minimum path

func Dijkstra(graph [][]int, start int) []int {
	n := len(graph)
	dist := make([]int, n) // 存储起点到每个节点的最短距离
	path := make([]int, n) // 最短路径的前顶点 如果这个距离较小 更改path中的值
	set := make([]bool, n) // 标记顶点是否纳入最短路径

	// 初始化数组
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
		path[i] = -1
	}
	//开始节点
	dist[start] = 0

	//Find the shortest dist and remark the index of top bit
	for i := 0; i < n-1; i++ {
		u := minDistance(dist, set)
		set[u] = true

		for v := 0; v < n; v++ {
			if !set[v] && graph[u][v] != 0 && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
				path[v] = u
			}
		}
	}
	return dist
}

// 找到未访问节点中具有最短距离的节点
func minDistance(dist []int, set []bool) int {
	min := math.MaxInt32
	minIndex := -1
	for v := 0; v < len(dist); v++ {
		if set[v] == false && dist[v] <= min {
			min = dist[v]
			minIndex = v
		}
	}
	return minIndex
}
