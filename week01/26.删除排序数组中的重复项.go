/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	len := 1
	val := nums[0]
	for _, v := range nums {
		if v == val {
			continue
		} else {
			index += 1
			nums[index] = v
			val = v
			len++
		}
	}
	return len
}

// @lc code=end

