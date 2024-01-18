/*
 * @Descripttion:
 * @Author:
 * @Date: 2024-01-17 22:41:37
 * @LastEditTime: 2024-01-17 22:41:49
 */
package heap

import "container/heap"

type Pair struct {
	i int
	j int
	v int
}

type MinHeap []Pair

func (this MinHeap) Len() int {
	return len(this)
}
func (this MinHeap) Less(i, j int) bool {
	return this[i].v < this[j].v
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

func kthSmallest(matrix [][]int, k int) int {
	mh := &MinHeap{}
	for index, value := range matrix[0] {
		// i横行 j竖列 v值  将第一个初始化堆
		heap.Push(mh, Pair{i: 0, j: index, v: value})
	}
	var cur Pair
	for !mh.Empty() && k > 0 {
		cur = heap.Pop(mh).(Pair)

		if cur.i < len(matrix)-1 {
			// 将下一行的值推入到堆中
			heap.Push(mh, Pair{i: cur.i + 1, j: cur.j, v: matrix[cur.i+1][cur.j]})

		}
		k--
	}
	return cur.v
}
