package heap

import "container/heap"

type PriorityQueue [][]int

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i][2] == pq[j][2] {
		return pq[i][1] > pq[j][1]
	}

	return pq[i][2] > pq[j][2]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	v := x.([]int)
	*pq = append(*pq, v)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func getSkyline(buildings [][]int) [][]int {
	pq := make(PriorityQueue, 0)
	var skyline [][]int
	i := 0
	nextEnd := 1000000
	curHeight := 0
	for i < len(buildings) || len(pq) > 0 {
		// 判断有没有建筑物需要处理
		if i < len(buildings) && buildings[i][0] <= nextEnd {
			curPoint := buildings[i][0]
			for i < len(buildings) && buildings[i][0] == curPoint {
				heap.Push(&pq, buildings[i])
				i++
			}

			// 最高建筑物的结束点
			nextEnd = pq[0][1]
			if pq[0][2] > curHeight {
				curHeight = pq[0][2]
				skyline = append(skyline, []int{curPoint, curHeight})
			}
		} else {
			// 优先级队列不为空时并且队列中最高的建筑物的结束点小于等于nextEnd的条件下,持续将最高建筑弹出
			for len(pq) > 0 && pq[0][1] <= nextEnd {
				heap.Pop(&pq)
			}
			if len(pq) > 0 {
				// 添加当前位置的天际线
				// pq[0][2] 是 heap中当前的最高的高度
				skyline = append(skyline, []int{nextEnd, pq[0][2]})
				curHeight = pq[0][2]
				nextEnd = pq[0][1] //结束点
			} else {
				skyline = append(skyline, []int{nextEnd, 0})
				curHeight = 0
				nextEnd = 100000
			}
		}
	}
	return skyline
}
