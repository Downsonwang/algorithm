/*
 * @Descripttion: 27 Remove Element - Array
 * @Author: DW
 * @Date: 2023-10-13 15:25:30
 * @LastEditTime: 2023-10-13 22:07:31
 */
package array

//暴力双循环
func removeElement(nums []int, val int) int {
	// nums [] all left right closed
	size := len(nums)

	for i := 0; i < size; i++ {
		if val == nums[i] {
			for j := i + 1; j < size; j++ {
				nums[j-1] = nums[j]
			}
			i--
			size--
		}
	}
	return size
}

// 没有改变元素的顺序
func removeElementBySlowandFastPointer(nums []int, val int) int {
	var slow int = 0
	var fast int

	for fast = 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// 改变元素的顺序和位置
func removeElementChangedBySlowandFastPointer(nums []int, val int) int {
	var left int = 0
	var right int = len(nums) - 1
	for left <= right {
		for left <= right && nums[left] != val {
			left++
		}
		for left <= right && nums[right] != val {
			right--
		}
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}
	return left
}
