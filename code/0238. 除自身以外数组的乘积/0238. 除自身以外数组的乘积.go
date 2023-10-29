package main

// https://leetcode.cn/problems/product-of-array-except-self/?envType=study-plan-v2&envId=top-interview-150
/*
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
*/

import (
	"fmt"
)

/*
遍历nums， 遍历到第i个时，可以通过左边 i-1个数， 和右边n-i个数的乘积
因此可以用一个left数组(长度n+1)存储所有在左边可能出现的乘积 left[0] = 1
可以用一个right数组(长度n+1)存储所有在right边可能出现的乘积 right[0] = 1
*/
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	left := make([]int, len(nums)+1)
	right := make([]int, len(nums)+1)
	left[0], right[0] = 1, 1
	// 分别计算所有左边可能出现的乘积
	left_sum := 1
	for i := 0; i < len(nums); i++ {
		left_sum = left_sum * nums[i]
		left[i+1] = left_sum
	}
	// 分别计算所有右边可能出现的乘积
	right_sum := 1
	for i := len(nums) - 1; i >= 0; i-- {
		right_sum = right_sum * nums[i]
		right[len(nums)-i] = right_sum
	}

	// 计算result结果
	for i := 0; i < len(nums); i++ {
		result[i] = left[i] * right[len(nums)-i-1]
	}
	return result

}
func main() {
	input := []int{1, 2, 3, 4}
	//target := longestPalindrome(str)
	result := productExceptSelf(input)
	fmt.Printf("result = %v", result)
}
