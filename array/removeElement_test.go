/*
 * @Descripttion:removeElements_test
 * @Author:DW
 * @Date: 2023-10-13 15:53:51
 * @LastEditTime: 2023-10-13 16:00:38
 */
package array

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	var nums = []int{3, 2, 2, 3}
	res := removeElement(nums, 3)
	fmt.Println(res)
}
