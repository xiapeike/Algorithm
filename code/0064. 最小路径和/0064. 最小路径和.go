package main

import (
	"code/util"
	"fmt"
)

// https://leetcode.cn/problems/minimum-path-sum/
/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

*/

/*
动态规划
dp[i][j] 表示达到该位置要走的路径
动态规划方程  dp[i][j] = min{dp[i-1][j] + dp[i][j-1]} + grid[i][j]
*/
func minPathSum064(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 0
		}
	}
	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = util.Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}
func main() {
	input := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	//input1 := [][]int{{1, 0}}
	result := minPathSum064(input)
	fmt.Printf("result = %v", result)
}
