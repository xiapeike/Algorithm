package main

import (
	"fmt"
)

// https://leetcode.cn/problems/climbing-stairs/
/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/

/*
递归
*/
func climbStairs070(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return climbStairs070(n-2) + climbStairs070(n-1)
}

/*
动规
*/
func climbStairs0702(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n-1]
}

func main() {
	input := 3
	result := climbStairs0702(input)
	fmt.Printf("result = %v", result)
}
