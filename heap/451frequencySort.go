package heap

import (
	"container/heap"
	"strings"
)

type KV struct {
	val       rune //  单词的索引
	frequency int  // 频次
}
type MaxHeap []KV

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].frequency > h[j].frequency }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(KV))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func frequencySort(s string) string {
	// strings.Builder 构建字符串类型
	// 提供了很多方法 用于向字符串中添加内容
	var result strings.Builder
	hp := &MaxHeap{}
	heap.Init(hp)
	hash := make(map[rune]int, 0)
	for _, value := range s {
		hash[value] = hash[value] + 1
	}
	for index, value := range hash {
		heap.Push(hp, KV{
			val:       index,
			frequency: value,
		})
	}
	//result.WriteRune(res.val): 在内部循环中，将当前元素的字符 val 写入字符串构建器 result 中，重复 res.frequency 次。
	// 这样就确保了按照字符频次构建的字符串
	for hp.Len() > 0 {
		res := heap.Pop(hp).(KV)
		for i := 0; i < res.frequency; i++ {
			result.WriteRune(res.val)
		}

	}
	return result.String()

}
