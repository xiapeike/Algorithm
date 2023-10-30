package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/jump-game-ii/
/*
给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
*/

func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		min_jump := math.MaxInt
		for j := 0; j < i; j++ {
			if dp[j] > -1 {
				if nums[j] >= i-j && dp[j]+1 < min_jump {
					min_jump = dp[j] + 1
					dp[i] = min_jump
				}
			}
		}

	}

	if dp[len(nums)-1] > -1 {
		return dp[len(nums)-1]
	}
	return 0
}

func main() {
	input := []int{2, 3, 0, 1, 4}
	result := jump(input)
	fmt.Printf("result = %v", result)
}
