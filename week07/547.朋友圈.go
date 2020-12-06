/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 朋友圈
 */

// @lc code=start
func findCircleNum(M [][]int) int {
	n := len(M)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	ans := n
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if M[i][j] == 1 {
				pi, pj := f(p, i), f(p, j)
				if pi != pj {
					p[pj] = pi
					ans--
				}
			}
		}
	}
	return ans
}
func f(p []int, x int) int {
	if p[x] != x {
		p[x] = f(p, p[x])
	}
	return p[x]
}

// @lc code=end

