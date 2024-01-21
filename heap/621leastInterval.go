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

func leastInterval(tasks []byte, n int) int {

	// 频率计算
	hash := map[byte]int{}
	// Record the frequency of  every task
	for _, v := range tasks {
		hash[v]++
	}

	// 维护一个最大堆
	hp := MaxHeap{}
	heap.Init(&hp)
	// 将每个任务的频率推进去
	for _, value := range hash {
		heap.Push(&hp, value)
	}
	// 记录所需总时间
	res := 0

	for hp.Len() > 0 {
		var tmp []int // 相同任务之间的冷却间隔
		k := n + 1    // 当前间隔大小 (+1 因为我们需要填充第一个任务) 确保第一个任务不被冷却

		for k > 0 && hp.Len() > 0 { // 处理当前间隔内的任务
			top := heap.Pop(&hp).(int)

			top-- // 减少任务的频率
			// ，每当执行一个任务后，间隔需要减小，

			k--                    //减小间隔大小
			res++                  // 增加总时间
			tmp = append(tmp, top) // 存储当前间隔内各个任务的频率信息。
		}

		for _, freq := range tmp {
			// 如果任务的频率大于零，则将其重新推送到任务堆中。
			if freq > 0 {
				heap.Push(&hp, freq)
			}
		}

		if hp.Len() == 0 {
			break
		}

		res += k
	}
	return res
}
