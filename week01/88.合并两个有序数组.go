/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	cap := m + n - 1
	index1 := m - 1
	index2 := n - 1
	for index1 >= 0 && index2 >= 0 {
		if nums1[index1] > nums2[index2] {
			nums1[cap] = nums1[index1]
			index1--
		} else {
			nums1[cap] = nums2[index2]
			index2--
		}
		cap--
	}
	if index1 < 0 {
		for {
			nums1[cap] = nums2[index2]
			cap--
			index2--
			if index2 < 0 {
				return
			}
		}
	}
	if index2 < 0 {
		for {
			nums1[cap] = nums1[index1]
			cap--
			index1--
			if index1 < 0 {
				return
			}
		}
	}
}

// @lc code=end

