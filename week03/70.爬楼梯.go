/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	tempDict := map[int]int{}
	return helper(n, tempDict)
}
func helper(n int, temp map[int]int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if val, ok := temp[n]; ok {
		return val
	} else {
		v := helper(n-1, temp) + helper(n-2, temp)
		temp[n] = v
		return v
	}
}

// @lc code=end

