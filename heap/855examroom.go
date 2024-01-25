/*
 * @Descripttion:
 * @Author:
 * @Date: 2024-01-25 11:03:57
 * @LastEditTime: 2024-01-25 15:14:07
 */
package heap

import "container/heap"

type seat struct {
	// 左右区间
	left, right int
}

// 第一个人 占 0 第二个人占 9 这时区间[1,8]第三个人就取4
func (s seat) size() int {
	return s.right - s.left
}

type ExamRoom struct {
	// index of the intervals containing key in the heap
	indices map[int]int
	h       []seat
	n       int
}

func (e ExamRoom) Len() int {
	return len(e.h)
}

func (e ExamRoom) Less(i, j int) bool {
	// 如果堆中的座位区间 相同
	if e.h[i].size()/2 == e.h[j].size()/2 {
		// we dont place elements in the middle

		if e.h[i].left == 0 && e.h[j].right == e.n-1 {
			return e.h[i].size() > e.h[j].size()
		}

		if e.h[j].left == 0 && e.h[i].right == e.n-1 {
			return e.h[i].size() > e.h[j].size()
		}
		// 	// 返回小号的
		return e.h[i].left < e.h[j].left
	}
	return e.h[i].size() > e.h[j].size()
}

func (e *ExamRoom) Swap(i, j int) {
	e.indices[e.h[i].left] = j
	e.indices[e.h[i].right] = j
	e.indices[e.h[j].left] = i
	e.indices[e.h[j].right] = i

	e.h[i], e.h[j] = e.h[j], e.h[i]
}

func (e *ExamRoom) Pop() interface{} {
	el := e.h[len(e.h)-1]
	// 弹出一个座位区间时,需要将这个座位区间的左右边界从索引映射中删除 他依不再堆中

	delete(e.indices, el.left)
	delete(e.indices, el.right)
	e.h = e.h[:len(e.h)-1]
	return el

}

func (e *ExamRoom) Push(v interface{}) {
	el := v.(seat)
	// 更新索引映射，将座位区间的左右边界与堆中的索引关联起来

	e.indices[el.left] = len(e.h)
	e.indices[el.right] = len(e.h)
	e.h = append(e.h, el)
}
func (e *ExamRoom) Remove(i int) *seat {
	if ind, ok := e.indices[i]; ok {
		last := len(e.h) - 1
		e.Swap(ind, last)
		v := e.Pop().(seat)
		if ind < len(e.h) {
			heap.Fix(e, ind)
		}
		return &v
	}
	return nil
}

func Constructor(n int) ExamRoom {
	ans := ExamRoom{
		indices: make(map[int]int),
		n:       n,
	}
	heap.Push(&ans, seat{0, n - 1})
	return ans
}

func (e *ExamRoom) Seat() int {
	best := heap.Pop(e).(seat)
	if best.left == 0 {
		if best.right >= 1 {
			heap.Push(e, seat{1, best.right})
		}
		return 0
	} else if best.right == e.n-1 {
		if best.right <= e.n-1 {
			heap.Push(e, seat{best.left, e.n - 2})
		}
		return e.n - 1
	} else {
		m := (best.left + best.right) / 2
		if best.left < m {
			heap.Push(e, seat{best.left, m - 1})
		}
		if best.right > m {
			heap.Push(e, seat{m + 1, best.right})
		}
		return m
	}
}

func (e *ExamRoom) Leave(p int) {
	left := e.Remove(p - 1)
	right := e.Remove(p + 1)
	if left != nil && right != nil {
		heap.Push(e, seat{left.left, right.right})
	} else if left != nil {
		heap.Push(e, seat{left.left, p})
	} else if right != nil {
		heap.Push(e, seat{p, right.right})
	} else {
		heap.Push(e, seat{p, p})
	}
}
