/*
 * @Descripttion:  stackbasedonqueue
 * @Author: DW
 * @Date: 2023-11-12 09:50:50
 * @LastEditTime: 2023-11-12 10:45:46
 */
package stack

type MyStack struct {
	QueueOne []int
	QueueTwo []int
}

func Constructor() MyStack {
	return MyStack{
		QueueOne: make([]int, 0),
		QueueTwo: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	oneLen, twoLen := len(this.QueueOne), len(this.QueueTwo)
	if oneLen != 0 {
		this.QueueTwo = append(this.QueueTwo, x)
		for i := 0; i < oneLen; i++ {
			this.QueueTwo = append(this.QueueTwo, this.QueueOne[i])
		}
		this.QueueOne = []int{}

	} else {
		this.QueueOne = append(this.QueueOne, x)
		for i := 0; i < twoLen; i++ {
			this.QueueOne = append(this.QueueOne, this.QueueTwo[i])
		}
		this.QueueTwo = []int{}
	}
}

func (this *MyStack) Pop() int {
	l1 := len(this.QueueOne)
	l2 := len(this.QueueTwo)

	var v int
	if l1 == 0 {
		// l2 里面取
		v = this.QueueTwo[0]
		this.QueueTwo = this.QueueTwo[1:]
	}

	if l2 == 0 {
		v = this.QueueOne[0]
		this.QueueOne = this.QueueOne[1:]
	}
	return v
}

func (this *MyStack) Top() int {
	l1 := len(this.QueueOne)
	l2 := len(this.QueueTwo)

	var v int
	if l1 == 0 {
		// l2 里面取
		v = this.QueueTwo[0]

	}

	if l2 == 0 {
		v = this.QueueOne[0]
	}
	return v
}

func (this *MyStack) Empty() bool {
	return len(this.QueueOne) == 0 && len(this.QueueTwo) == 0
}
