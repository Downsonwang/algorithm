/*
 * @Descripttion: LinkedList CRUD
 * @Author:DW
 * @Date: 2023-10-29 20:26:31
 * @LastEditTime: 2023-10-29 22:17:22
 */
package linkedlist

import "fmt"

// One node
type LinkedNode struct {
	next  *LinkedNode
	value int
}

// All LinkedList
type LinkedList struct {
	head   *LinkedNode
	length int
}

func NewListNode(v int) *LinkedNode {
	return &LinkedNode{nil, v}
}

// getNext() .
func (this *LinkedNode) GetNext() *LinkedNode {
	return this.next
}

// getValue() .
func (this *LinkedNode) GetValue() int {
	return this.value
}

// NewLinkedList() .
func NewLinkedList() *LinkedList {
	return &LinkedList{head: NewListNode(0), length: 0}
}

// InsertValueAfterNode() .
func (this *LinkedList) InsertAfter(p *LinkedNode, val int) bool {
	if p == nil {
		return false
	}
	new_node := NewListNode(val)
	oldNext := p.next
	p.next = new_node
	new_node.next = oldNext
	this.length++
	return true
}

// InsertValueBeforeNode() 。
func (this *LinkedList) InsertBefore(p *LinkedNode, val int) bool {
	// alert it !!!
	if p == nil || p == this.head {
		return false
	}
	cur := this.head.next
	prev := this.head
	if cur == nil {
		return false
	}
	// p is not exist in the list
	for cur != nil {
		if cur == p {
			break
		}
		prev = cur
		cur = cur.next
	}

	new_node := NewListNode(val)
	prev.next = new_node
	new_node.next = cur
	this.length++
	return true
}

// InsertToHead()
func (this *LinkedList) InsertToHead(val int) bool {
	return this.InsertAfter(this.head, val)
}

// InsertToTail()
func (this *LinkedList) InsertToTail(val int) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	return this.InsertAfter(cur, val)
}

// GetNodeByIndex 通过索引开始找
func (this *LinkedList) GetNodeByIndex(index int) *LinkedNode {
	if index >= this.length {
		return nil
	}
	// head node doesnt contain data so cur := this.head.next
	cur := this.head.next
	for i := 0; i < index; i++ {
		cur = cur.next

	}
	return cur
}

//DeleteArgumentNode
func (this *LinkedList) DeleteNode(p *LinkedNode) bool {
	if p == nil {
		return false
	}
	cur := this.head.next
	prev := this.head
	for cur != nil {
		if cur == p {
			break
		}

		prev = cur
		cur = cur.next
		return false
	}

	if cur == nil {
		return false
	}
	prev.next = p.next
	// 释放内存(防止内存泄漏). p=nil
	p = nil
	this.length--
	return true
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
