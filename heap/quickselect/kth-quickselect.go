/*
 * @Descripttion: QuickSelect
 * @Author:Downson
 * @Date: 2024-01-07 11:26:22
 * @LastEditTime: 2024-01-07 20:39:16
 */
package quickselect

import "math/rand"

// QuickSort
func findKthLargest(nums []int, k int) int {
	low, high := 0, len(nums)-1

	// target location
	target := len(nums) - k
	var cur int
	// 缩小分区
	for low <= high {
		cur = partition(nums, low, high)
		if cur == target {
			break
		}
		if cur < target {
			low = cur + 1
		} else {
			high = cur - 1
		}
	}
	return nums[cur]

}

func partition(nums []int, low, high int) int {
	if low == high {
		return low
	}
	randIdx := low + rand.Intn(high-low+1)
	num := nums[randIdx]
	swap(nums, randIdx, high)

	left, right := low, high-1
	for left <= right {
		for left <= right && nums[left] <= num {
			left++
		}

		for left <= right && nums[right] > num {
			right--
		}
		if left < right {
			swap(nums, left, right)
		}
	}
	swap(nums, left, high)

	return left

}

func swap(nums []int, left, right int) {
	nums[left], nums[right] = nums[right], nums[left]
}
