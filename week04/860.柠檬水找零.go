/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	count5 := 0
	count10 := 0
	for _, bill := range bills {
		if bill == 5 {
			count5++
		} else if bill == 10 {
			count10++
			count5--
		} else if bill == 20 {
			if count10 > 0 && count5 > 0 {
				count10--
				count5--
			} else {
				count5 = count5 - 3
				if count5 < 0 {
					return false
				}
			}
		}
	}
	return count10 >= 0 && count5 >= 0
}

// @lc code=end

