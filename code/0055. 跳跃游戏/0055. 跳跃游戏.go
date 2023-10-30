package main

import (
	"code/util"
	"fmt"
)

// https://leetcode.cn/problems/jump-game/
/*
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。

*/

// 动态规划法 dp[n] bool
// dp[i]表示是否能从起点跳到i这个位置
// dp[i] 的取值由dp[0]~dp[i-1]为true的点，是否能达到dp[i]决定
func canJump055(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && nums[j] >= i-j {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(nums)-1]
}

/*
贪心算法
遍历数组，判断每一次覆盖范围是否会变大，直到覆盖到最后一位
*/
func canJump0552(nums []int) bool {
	cover_range := 0

	for i := 0; i < len(nums); i++ {
		if i <= cover_range {
			cover_range = util.Max(nums[i]+i, cover_range)
		}
	}

	if cover_range >= len(nums)-1 {
		return true
	}

	return false
}

func main() {
	input := []int{2, 3, 1, 1, 4}
	result := canJump0552(input)
	fmt.Printf("result = %v", result)
}
