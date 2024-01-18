/*
 * @Descripttion: ugly number 2
 * @Author: DW
 * @Date: 2024-01-15 21:27:43
 * @LastEditTime: 2024-01-15 21:46:16
 */
package heap

import "container/heap"

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

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}

	hp := &IntHeap{1}
	heap.Init(hp)

	primes := []int{2, 3, 5}
	var seen map[int]struct{}
	seen[1] = struct{}{}

	for i := 1; i < n; i++ {
		cur := heap.Pop(hp).(int)
		for _, val := range primes {
			next := cur * val
			if _, found := seen[next]; !found {
				heap.Push(hp, next)
				seen[next] = struct{}{}
			}
		}
	}
	return heap.Pop(hp).(int)
}
