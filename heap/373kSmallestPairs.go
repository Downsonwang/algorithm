/*
 * @Descripttion: Find K pairs with smallest sums
 * @Author: DW
 * @Date: 2024-01-17 20:48:54
 * @LastEditTime: 2024-01-17 21:47:15
 */
package heap

import "container/heap"

type Pair struct {
	u   int //index of nums1
	v   int // index of nums2
	sum int
}

type MinHeap []Pair

func (this MinHeap) Len() int {
	return len(this)
}
func (this MinHeap) Less(i, j int) bool {
	return this[i].sum < this[j].sum
}
func (this MinHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func (this MinHeap) Empty() bool {
	return len(this) == 0
}
func (this *MinHeap) Push(x interface{}) {
	*this = append(*this, x.(Pair))
}
func (this *MinHeap) Pop() interface{} {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[:n-1]
	return x

}
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := make([][]int, 0)
	// 0. 维护一个最小堆
	minheap := &MinHeap{}
	heap.Init(minheap)

	// 1. 遍历nums1 和 nums2 相加 并记录当前加的索引,
	// 2. 将相加的结果Push到最小堆
	for index, v1 := range nums1 {
		heap.Push(minheap, Pair{u: index, v: 0, sum: v1 + nums2[0]})
	}
	// 3. 假如这个堆的长度为k Pop出前k个值
	for !minheap.Empty() && k > 0 {
		curMin := heap.Pop(minheap).(Pair) // 拿到当前的最小
		// // 将其相加值的u v 拿到
		l := curMin.u
		r := curMin.v
		// 5. 添加到最后的output中
		result = append(result, []int{nums1[l], nums2[r]})

		// 6. 遍历nums2
		nextIdx := r + 1
		if nextIdx < len(nums2) {
			//如果对于当前的nums1元素在nums2中还有更多元素，则将下一个对推入最小堆。
			heap.Push(minheap, Pair{u: l, v: nextIdx, sum: nums1[l] + nums2[nextIdx]})
		}
		k--

	}

	// 6. 输出结果
	return result
}
