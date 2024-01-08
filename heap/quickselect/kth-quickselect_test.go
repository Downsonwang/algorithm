/*
 * @Descripttion: TestFindKthLargest
 * @Author:DW
 * @Date: 2024-01-07 14:31:11
 * @LastEditTime: 2024-01-07 14:44:45
 */
package quickselect

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2},
		// Add more test cases as needed
	}

	for _, test := range tests {
		result := FindKthLargest(test.nums, test.k)
		// You can add any additional assertions or logging if needed
		if result < 0 {
			t.Errorf("Negative result for nums=%v, k=%d: %d", test.nums, test.k, result)
		}
	}

}
