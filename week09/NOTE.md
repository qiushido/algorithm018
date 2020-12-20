学习笔记

不同路径ii 状态转移方程
dp[i][j] = 0 (i,j == 0)
dp[i][j] = dp[i][j-1] + dp[i-1][j] (i,j != 0)
