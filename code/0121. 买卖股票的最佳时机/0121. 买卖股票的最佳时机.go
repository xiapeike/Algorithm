package main

import (
	"fmt"
)

/*
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&envId=top-interview-150

给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0
*/

/*
遍历数组，分别判断在当天买入 后续卖出的最大收益
时间复杂度O(n^2)
*/
func maxProfit(prices []int) int {
	result := 0
	if len(prices) <= 1 {
		return result
	}
	for i := 0; i < len(prices)-1; i++ {
		item_max := 0
		for j := i + 1; j < len(prices); j++ {
			if prices[j] > prices[i] && prices[j] > item_max {
				item_max = prices[j]
			}
		}
		if item_max-prices[i] > result {
			result = item_max - prices[i]
		}
	}
	return result
}

/*
动态规划思想
用一个数组dp[i] 记录前i天的最低值
当遍历i+1天时，只需要判断第i+1天的价格和前i天的最低值的差价，即可得到最终差价
*/
func maxProfit1(prices []int) int {
	result := 0
	if len(prices) <= 1 {
		return result
	}
	// dp[i] 记录前i天的最低值
	dp := make([]int, len(prices))
	dp[0] = prices[1]
	for i := 1; i < len(prices); i++ {
		if prices[i] < dp[i-1] {
			dp[i] = prices[i]
		} else {
			dp[i] = dp[i-1]
			if prices[i]-dp[i-1] > result {
				result = prices[i] - dp[i-1]
			}
		}

	}
	return result
}

/*
针对方法2，其实不需要一个dp数组来记录前i的最低值
只需要用一个单独的变量来记录即可
*/
func maxProfit2(prices []int) int {
	result := 0
	if len(prices) <= 1 {
		return result
	}
	// dp[i] 记录前i天的最低值
	cost := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < cost {
			cost = prices[i]
		} else {
			if prices[i]-cost > result {
				result = prices[i] - cost
			}
		}

	}
	return result
}
func main() {
	input := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit2(input)
	fmt.Printf("sum = %+v", result)
}
