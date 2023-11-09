/*
 * @Descripttion:24. swapPairs
 * @Author: DW
 * @Date: 2023-11-08 17:53:55
 * @LastEditTime: 2023-11-08 18:58:53
 */
package linkedlist

func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prev := &ListNode{Val: 0, Next: head}
	var cur *ListNode = prev
	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next
		tmp1 := cur.Next.Next.Next
		cur.Next = cur.Next.Next

		cur.Next.Next = tmp
		cur.Next.Next.Next = tmp1

		cur = cur.Next.Next
	}
	return prev.Next
}
