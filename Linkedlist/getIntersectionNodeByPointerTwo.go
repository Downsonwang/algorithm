/*
 * @Descripttion: getIntersectionNodeByPointerTwo
 * @Author:DW
 * @Date: 2023-11-09 19:05:08
 * @LastEditTime: 2023-11-09 19:30:31
 */
package linkedlist

func getIntersectionNodeByPointerTwo(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB
	// 双指针 跟相同的路径跑步偶遇类型
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headB
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = headA
		}
	}
	return l1
}
