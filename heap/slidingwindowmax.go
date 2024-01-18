package heap

import "container/heap"

// Heap Priority Queue
type KV struct {
	value, index int
}
type MaxHeap []KV

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].value > h[j].value }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(KV))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func maxSlidingWindow(nums []int, k int) []int {
	if k <= 0 || len(nums) == 0 {
		return nil
	}

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	res := make([]int, len(nums)-k+1)

	// i为窗口开始位置 j为窗口结束位置 index 结果切片的索引
	left, right, index := 0, 0, 0

	for right < len(nums) {
		kv := KV{right, nums[right]}
		heap.Push(maxHeap, kv)
		if right-left+1 < k {
			right++
		} else if right-left+1 == k {
			// 最大值的索引比当前窗口的索引还小
			for (*maxHeap)[0].index < left {
				// 解散这个堆
				heap.Pop(maxHeap)
			}
			first := (*maxHeap)[0]
			res[index] = first.value
			// 没在下一个位置
			if first.index < left+1 {
				heap.Pop(maxHeap)
			}
			left++
			right++
			index++
		}
	}
	return res
}
