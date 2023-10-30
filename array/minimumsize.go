/*
 * @Descripttion:
 * @Author:DW

 * @Date: 2023-10-22 21:52:34
 * @LastEditTime: 2023-10-22 21:53:42
 */
package array

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	slow := 0
	fast := slow
	min_sum := make([]int, n)
	for slow < n {
		if target-nums[slow] != nums[fast] {
			fast++
		} else {
			min := fast - slow
			min_sum = append(min_sum, min)

		}
		slow++
	}

	min_data := min_sum[0]
	for _, v := range min_sum {
		if v < min_data {
			min_data = v
		}

	}

	return min_data

}
