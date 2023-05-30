package main

import (
	"fmt"
)

// https://leetcode.cn/problems/unique-paths-ii/
/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。

*/

/*
动态规划
如果有障碍，dp[i][j] = 0
动态规划方程  dp[i][j] = dp[i-1][j] + dp[i][j-1]
*/
func uniquePathsWithObstacles063(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 0
		}
	}

	if obstacleGrid[0][0] != 1 {
		dp[0][0] = 1
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] != 1 {
			dp[i][0] = dp[i-1][0]
		}
	}

	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] != 1 {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
func main() {
	input := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	//input1 := [][]int{{1, 0}}
	result := uniquePathsWithObstacles063(input)
	fmt.Printf("result = %v", result)
}
