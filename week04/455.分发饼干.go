/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ans := 0
	for _, v := range s {
		if ans == len(g) {
			break
		}
		if g[ans] <= v {
			ans++
		}
	}
	return ans
}

// @lc code=end

