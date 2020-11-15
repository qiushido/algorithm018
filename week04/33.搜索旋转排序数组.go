/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[lo] {
			if target < nums[mid] && target >= nums[lo] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

// @lc code=end

