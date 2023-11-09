/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-11-08 18:41:12
 * @LastEditTime: 2023-11-08 18:47:19
 */
package linkedlist

import (
	"fmt"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	res := SwapPairs(head)
	fmt.Println(res)
}
