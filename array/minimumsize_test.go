/*
 * @Descripttion:
 * @Author:DW
 * @Date: 2023-10-22 21:53:51
 * @LastEditTime: 2023-10-22 21:56:49
 */
package array

import (
	"fmt"
	"testing"
)

func TestMinimumsize(t *testing.T) {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	res := minSubArrayLen(target, nums)
	fmt.Println(res)
}
