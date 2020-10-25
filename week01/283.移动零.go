/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	flag := 0
	for i, v := range nums {
		if v == 0 {
			continue
		} else {
			nums[i] = 0
			nums[flag] = v
			flag++
		}
	}
}

// @lc code=end

