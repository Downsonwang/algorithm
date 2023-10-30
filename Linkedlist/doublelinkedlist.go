/*
 * @Descripttion: DoubleLinkedList
 * @Author:DW
 * @Date: 2023-10-30 09:46:37
 * @LastEditTime: 2023-10-30 11:40:13
 */
package linkedlist

type DoubleLinkedNode struct {
	prev_node *DoubleLinkedNode
	next_node *DoubleLinkedNode
	val       int
}

type DoubleLinkedList struct {
	head   *DoubleLinkedNode
	length int
}

func NewListNode(v int) *DoubleLinkedNode {
	return &DoubleLinkedNode{nil, nil, v}
}

func (this *DoubleLinkedList) IsEmpty() bool {
	if this.head == nil {
		return true
	}
	return false
}

func (this *DoubleLinkedList) Length() int {
	if this.IsEmpty() {
		return 0
	}
	count := 0
	cur := this.head.next_node

	if cur != nil {
		count++
		cur = cur.next_node
	}
	return count

}

func (this *DoubleLinkedList) AddFromHead(v int) bool {
	new_node := &DoubleLinkedNode{val: v}
	if this.IsEmpty() {
		this.head = new_node
		return true
	}
	this.head.prev_node = new_node
	new_node.next_node = this.head
	this.head = new_node
	return true
}

func (this *DoubleLinkedList) AddFromTail(v int) bool {
	new_node := &DoubleLinkedNode{val: v}
	if this.IsEmpty() {
		this.head = new_node
		return true
	}
	cur := this.head.next_node
	for cur.next_node != nil {
		cur = cur.next_node
	}
	cur.next_node = new_node
	new_node.prev_node = cur
	return true
}

func (this *DoubleLinkedList) InsertAfter(node *DoubleLinkedNode, val int) bool {
	if node == nil {
		return false
	}

	new_node := NewListNode(val)

	// 更新新节点的指针
	new_node.prev_node = node
	new_node.next_node = node.next_node
	// 更新原节点的下一个结点的指针
	if node.next_node != nil {
		node.next_node.prev_node = new_node
	}
	// 更新原节点的指针
	node.next_node = new_node
	this.length++
	return true
}

func (this *DoubleLinkedList) InsertBefore(node *DoubleLinkedNode, val int) bool {
	if node == nil {
		return false
	}

	new_node := NewListNode(val)

	// 更新新节点的指针
	new_node.prev_node = node.prev_node
	new_node.next_node = node

	// 更新原节点的前一个结点的指针
	if node.prev_node != nil {
		node.prev_node.next_node = new_node
	} else {
		this.head = new_node //如果P是头结点,更新链表头.
	}
	// 更新原节点的指针
	node.prev_node = new_node
	this.length++

	return true
}

// DeleteHead .
func (this *DoubleLinkedList) DeleteHead() bool {
	if this.IsEmpty() {
		return false //链表为空 删除
	}

	// 获取原节点
	oldHead := this.head

	// 更新链表的头结点
	this.head = oldHead.next_node

	// 将新头节点的prev指针为nil
	if this.head != nil {
		this.head.prev_node = nil
	}
	//释放原头结点
	oldHead = nil
	this.length--
	return true
}

// DeleteTail .
func (this *DoubleLinkedList) DeleteTail() bool {
	if this.IsEmpty() {
		return false
	}

	cur := this.head

	for cur.next_node != nil {
		cur = cur.next_node
	}

	//删除尾节点
	if cur.prev_node != nil {
		cur.prev_node.next_node = nil
	} else {
		// 链表只有一个结点
		this.head = nil

	}
	cur = nil
	this.length--
	return true
}
