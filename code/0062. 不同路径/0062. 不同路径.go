package main

import (
	"fmt"
)

// https://leetcode.cn/problems/unique-paths/
/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

*/

/*
动态规划
我们令 dp[i][j] 是到达 i, j 最多路径

动态方程：dp[i][j] = dp[i-1][j] + dp[i][j-1]
*/
func uniquePaths062(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 0
		}
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	// dp[i][j] = dp[i-1][j] + dp[i][j-1]
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func main() {
	result := uniquePaths062(3, 7)
	fmt.Printf("result = %v", result)
}
