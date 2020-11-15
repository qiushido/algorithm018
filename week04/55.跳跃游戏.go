/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	farest := 0
	for i, v := range nums {
		if farest >= i && i+v > farest {
			farest = i + v
		}
	}
	return farest >= len(nums)-1
}

// @lc code=end

