/*
 * @Descripttion: leetcode - 215kth-largest-element-in-an-array - Heapify
 * @Author:DW
 * @Date: 2024-01-06 17:45:19
 * @LastEditTime: 2024-01-07 10:28:47
 */
package heap

import (
	"container/heap"
	"unsafe"
)

// O(log n k)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func findKthLargest(nums []int, k int) int {
	h := (*IntHeap)(unsafe.Pointer(&nums))
	heap.Init(h)
	for h.Len() != k {
		heap.Pop(h)
	}
	return heap.Pop(h).(int)
}
