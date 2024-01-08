package quickselect

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type IntHeap []*ListNode

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	ptr := head
	h := &IntHeap{}

	for _, list := range lists {
		if list != nil {
			*h = append(*h, list)
		}

	}

	heap.Init(h)

	for h.Len() > 0 {
		smaller := heap.Pop(h).(*ListNode)
		ptr.Next = smaller
		smaller = smaller.Next
		if smaller != nil {
			heap.Push(h, smaller)
		}
		ptr = ptr.Next
	}
	return head.Next
}
