/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	temp := make([]int, 26)
	for i, v := range s {
		temp[v-'a']++
		temp[t[i]-'a']--
	}
	for _, v := range temp {
		if v != 0 {
			return false
		}
	}
	return true
}

// @lc code=end

