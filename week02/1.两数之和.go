/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	tempHash := map[int]int{}
	for i, v := range nums {
		if p, ok := tempHash[target-v]; ok {
			return []int{p, i}
		}
		tempHash[v] = i
	}
	return []int{}
}

// @lc code=end

