/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	queue := []*Node{}
	// 入队
	queue = append(queue, root)
	for len(queue) > 0 {
		temp := []int{}
		i := len(queue)
		for m := 0; m < i; m++ {
			// 出队
			node := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			temp = append(temp, node.Val)
			for _, v := range node.Children {
				queue = append([]*Node{v}, queue...)
			}
		}
		ans = append(ans, temp)
	}
	return ans
}

// @lc code=end

