/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
func reverseWords(s string) string {
	res := ""
	temp := ""
	for _, v := range s {
		vv := string(v)
		if vv == " " && temp != "" && res == "" {
			res = temp
		} else if vv == " " && temp != "" && res != "" {
			res = temp + " " + res
		}
		if vv == " " {
			temp = ""
		} else {
			temp = temp + vv
		}
	}
	if temp == "" {
		return strings.Trim(res, " ")
	}
	return strings.Trim(temp+" "+res, " ")
}

// @lc code=end

