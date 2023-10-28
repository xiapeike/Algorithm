package main

import (
	"fmt"
)

/*
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/?envType=study-plan-v2&envId=top-interview-150
给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。

返回 你能获得的 最大 利润 。
*/

/*
将每一段的上升的值都加在一起
*/
func maxProfit(prices []int) int {
	result := 0
	if len(prices) <= 1 {
		return result
	}
	cost := prices[0]
	sub_sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			// 依然处于上升段
			sub_sum = prices[i] - cost
		} else {
			result = result + sub_sum
			sub_sum = 0
			cost = prices[i]
		}

	}
	if sub_sum > 0 {
		result += sub_sum
	}
	return result
}

func main() {
	input := []int{1, 2, 3, 4, 5}
	result := maxProfit(input)
	fmt.Printf("sum = %+v", result)
}
