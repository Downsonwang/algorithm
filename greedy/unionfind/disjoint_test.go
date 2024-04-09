/*
 * @Descripttion:  TEst _ 并查集
 * @Author: DW
 * @Date: 2024-04-08 22:17:02
 * @LastEditTime: 2024-04-08 22:19:46
 */
package unionfind

import (
	"fmt"
	"testing"
)

func TestFindConnectedComponents(t *testing.T) {
	graph := map[int][]int{
		0: {1, 2},
		1: {0, 2},
		2: {0, 1},
		3: {4},
		4: {3},
	}

	connectedComponents := FindConnectedComponents(graph)
	fmt.Println(connectedComponents)
}
