/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	tempHash := make(map[[26]int][]string)
	for _, v := range strs {
		if _, ok := tempHash[makeKey(v)]; ok {
			tempHash[makeKey(v)] = append(tempHash[makeKey(v)], v)
		} else {
			val := make([]string, 0)
			tempHash[makeKey(v)] = append(val, v)
		}
	}
	ans := make([][]string, 0, len(tempHash))
	for _, v := range tempHash {
		ans = append(ans, v)
	}
	return ans
}

func makeKey(s string) [26]int {
	listInt := [26]int{}
	for _, v := range s {
		listInt[v-'a']++
	}
	return listInt
}

// @lc code=end

