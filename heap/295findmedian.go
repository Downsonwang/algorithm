/*
 * @Descripttion: Heap - 295 Find Median from Data stream
 * @Author: DW
 * @Date: 2024-01-16 09:00:18
 * @LastEditTime: 2024-01-16 10:46:46
 */
package heap

import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Minheap []int

func (h Minheap) Len() int           { return len(h) }
func (h Minheap) Less(i, j int) bool { return h[i] < h[j] }
func (h Minheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Minheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Minheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 数据结构 - ordered int list
type MedianFinder struct {
	minHeap Minheap
	maxHeap MaxHeap
}

// 初始化 一个大顶堆 和 一个小顶堆
func Constructor() MedianFinder {
	return MedianFinder{minHeap: nil, maxHeap: nil}
}

// 假如当前num比  大顶堆倒着 小顶堆放大值 这样就会找到中间值
func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == 0 {
		heap.Push(&this.maxHeap, num)
	} else if num > this.maxHeap[0] {
		heap.Push(&this.minHeap, num)
	} else {
		heap.Push(&this.maxHeap, num)
	}

	if this.minHeap.Len()+1 < this.maxHeap.Len() {
		heap.Push(&this.minHeap, heap.Pop(&this.maxHeap))
	} else if this.maxHeap.Len() < this.minHeap.Len() {
		heap.Push(&this.maxHeap, heap.Pop(&this.minHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	maxSize := this.maxHeap.Len()
	minSize := this.minHeap.Len()
	allSize := maxSize + minSize
	res := 0.0
	if allSize%2 == 0 {
		res = float64(this.maxHeap[0]+this.minHeap[0]) / 2.0
	} else if maxSize > minSize {

		res = float64(this.maxHeap[0])

	} else if minSize > maxSize {
		res = float64(this.minHeap[0])
	}
	return res
}
