/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-11-07 21:08:23
 * @LastEditTime: 2023-11-07 22:30:29
 */
package stack

/**
   # 数组实现顺序栈
**/

type Arr struct {
	items []int
	top   int
}

func NewArr() *Arr {
	return &Arr{items: make([]int, 0), top: -1}
}

func (this *Arr) IsEmpty() bool {
	if this.top < 0 {
		return true
	}
	return false
}

func (this *Arr) Push(val int) {
	if this.top < 0 {
		this.top = 0
	} else {
		this.top += 1
	}

	if this.top > len(this.items)-1 {
		this.items = append(this.items, val)
	} else {
		this.items[this.top] = val
	}

}

func (this *Arr) Pop() int {
	if this.IsEmpty() {
		return -1
	}
	v := this.items[this.top]
	this.top -= 1
	return v
}

func (this *Arr) Top() int {
	if this.IsEmpty() {
		return -1
	}
	v := this.items[this.top]
	return v
}

func (this *Arr) Flush() {
	this.top = -1
}
