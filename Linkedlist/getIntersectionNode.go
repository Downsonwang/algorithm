/*
 * @Descripttion:  getIntersectionNode
 * @Author:  DW
 * @Date: 2023-11-09 14:06:36
 * @LastEditTime: 2023-11-09 14:21:15
 */
package linkedlist

// 交点不是数值相同 而是指针相同
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 循环遍历这两个列表
	curA := headA
	curB := headB
	lenA := 0
	lenB := 0
	for curA != nil {
		lenA++
		curA = curA.Next
	}

	for curB != nil {
		lenB++
		curB = curB.Next
	}
	// 计算出链表的差值,让多的先移动gap位 和 少的对齐
	var gap int
	var slow, fast *ListNode
	if lenA > lenB {
		gap = lenA - lenB
		fast, slow = headA, headB
	} else {
		gap = lenB - lenA
		fast, slow = headB, headA
	}
	for i := 0; i <= gap; i++ {
		fast = fast.Next
	}
	// 找是否相同
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	// 相同则跳出遍历
	// 不同空指针
	return fast

}
