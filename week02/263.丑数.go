/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	return search(num)
}

func search(num int) bool {
	if num == 1 {
		return true
	}
	if num%2 != 0 && num%3 != 0 && num%5 != 0 {
		return false
	} else {
		var f1 bool
		if (num % 2) == 0 {
			f1 = search(num / 2)
		} else {
			f1 = false
		}
		var f2 bool
		if (num % 3) == 0 {
			f2 = search(num / 3)
		} else {
			f2 = false
		}
		var f3 bool
		if (num % 5) == 0 {
			f3 = search(num / 5)
		} else {
			f3 = false
		}
		return (f1 || f2 || f3)
	}
}

// @lc code=end

