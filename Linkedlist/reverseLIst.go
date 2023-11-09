/*
 * @Descripttion:
 * @Author:DW
 * @Date: 2023-11-08 17:28:12
 * @LastEditTime: 2023-11-08 17:52:19
 */
package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	cur := head
	for cur != nil {
		// 标记next
		var next *ListNode = cur.Next
		// 指针反转
		cur.Next = prev
		// 移动
		prev = cur
		// 移动
		cur = next

	}
	return cur
}
