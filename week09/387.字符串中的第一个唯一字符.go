/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	var tmp [26]int
	for i, v := range s {
		tmp[v-'a'] = i
	}
	for i, v := range s {
		if tmp[v-'a'] == i {
			return i
		} else {
			tmp[v-'a'] = -1
		}
	}
	return -1
}

// @lc code=end

