/*
 * @Descripttion: 链式栈
 * @Author:DW
 * @Date: 2023-11-07 22:34:17
 * @LastEditTime: 2023-11-07 22:47:15
 */
package stack

type node struct {
	next *node
	val  int
}

type LinkedListStack struct {
	topNode *node // 栈顶节点
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		nil,
	}
}

func (this *LinkedListStack) IsEmpty() bool {
	return this.topNode == nil
}

func (this *LinkedListStack) Push(val int) {
	this.topNode = &node{next: this.topNode, val: val}
}

func (this *LinkedListStack) Pop() int {
	if this.IsEmpty() {
		return -1
	}
	v := this.topNode.val
	this.topNode = this.topNode.next
	return v

}

func (this *LinkedListStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.topNode.val
}

func (this *LinkedListStack) Flush() {
	this.topNode = nil
}
