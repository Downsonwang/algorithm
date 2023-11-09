/*
 * @Descripttion:19
 * @Author:DW
 * @Date: 2023-11-08 19:29:23
 * @LastEditTime: 2023-11-09 13:01:52
 */
package linkedlist

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Val: 999, Next: head}

	slow, fast := dummyHead, dummyHead

	for i := n; i > 0; i-- {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dummyHead.Next
}
