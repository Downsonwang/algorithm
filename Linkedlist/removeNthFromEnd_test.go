/*
 * @Descripttion:19 Remove Nth Node From End of List
 * @Author: DW
 * @Date: 2023-11-09 09:55:42
 * @LastEditTime: 2023-11-09 09:59:45
 */
package linkedlist

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	n := 2
	res := RemoveNthFromEnd(head, n)
	fmt.Println(res)
}
