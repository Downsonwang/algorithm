/*
 * @Descripttion: 938Desc
 * @Author: DW
 * @Date: 2024-01-08 10:48:48
 * @LastEditTime: 2024-01-08 11:31:57
 */
package heap

import (
	"container/heap"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Item struct {
	Node *TreeNode
}

type IntHeap []*Item

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].Node.Val < h[j].Node.Val }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	pq := make(IntHeap, 0)
	heap.Init(&pq)
	sum := 0
	heap.Push(&pq, &Item{Node: root})
	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		node := item.Node

		if node.Val >= low && node.Val <= high {
			sum += node.Val
		}

		if node.Left != nil && node.Val > low {
			heap.Push(&pq, &Item{Node: node.Left})
		}

		if node.Right != nil && node.Val < high {
			heap.Push(&pq, &Item{Node: node.Right})
		}
	}
	return sum
}
