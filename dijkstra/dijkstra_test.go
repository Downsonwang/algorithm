/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-10-31 11:26:43
 * @LastEditTime: 2023-10-31 11:46:13
 */
package dijkstra

import (
	"fmt"
	"testing"
)

func TestDijkstra(t *testing.T) {
	graph := [][]int{
		{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{4, 0, 8, 0, 0, 0, 0, 11, 0},
		{0, 8, 0, 7, 0, 4, 0, 0, 2},
		{0, 0, 7, 0, 9, 14, 0, 0, 0},
		{0, 0, 0, 9, 0, 10, 0, 0, 0},
		{0, 0, 4, 14, 10, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 1, 6},
		{8, 11, 0, 0, 0, 0, 1, 0, 7},
		{0, 0, 2, 0, 0, 0, 6, 7, 0},
	}

	startNode := 0
	distances := Dijkstra(graph, startNode)

	fmt.Println("Shortest distances from node", startNode+1, ":")
	for i := 0; i < len(distances); i++ {
		fmt.Println("Node", i+1, ":", distances[i])
	}
}
