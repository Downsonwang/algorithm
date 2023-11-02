/*
 * @Descripttion: 707-Design Linked List
 * @Author:DW
 * @Date: 2023-11-01 20:04:44
 * @LastEditTime: 2023-11-02 00:21:19
 */
package linkedlist

type MyLinkedNode struct {
	val  int
	next *MyLinkedNode
}
type MyLinkedList struct {
	head *MyLinkedNode
	len  int
}

func Constructor() MyLinkedList {
	newNode := &MyLinkedNode{
		-999,
		nil,
	}
	return MyLinkedList{head: newNode, len: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.len || this == nil {
		return -1
	}
	cur := this.head.next

	for i := 0; i < index; i++ {
		cur = cur.next

	}
	return cur.val

}

func (this *MyLinkedList) AddAtHead(val int) {
	new_node := &MyLinkedNode{val: val}
	// 新节点
	new_node.next = this.head.next
	this.head.next = new_node
	this.len++
}

func (this *MyLinkedList) AddAtTail(val int) {
	new_node := &MyLinkedNode{val: val}

	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = new_node
	this.len++

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		index = 0
	} else if index > this.len {
		return
	}
	new_node := &MyLinkedNode{val: val, next: nil}
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.next

	}

	new_node.next = cur.next
	cur.next = new_node
	this.len++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.len || index < 0 {
		return
	}
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	if cur.next != nil {
		cur.next = cur.next.next
	}

	this.len--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
