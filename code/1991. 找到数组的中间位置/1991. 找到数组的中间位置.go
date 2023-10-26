package main

import "fmt"

/*
https://leetcode.cn/problems/find-the-middle-index-in-array/
给你一个下标从 0 开始的整数数组 nums ，请你找到 最左边 的中间位置 middleIndex （也就是所有可能中间位置下标最小的一个）。

中间位置 middleIndex 是满足 nums[0] + nums[1] + ... + nums[middleIndex-1] == nums[middleIndex+1] + nums[middleIndex+2] + ... + nums[nums.length-1] 的数组下标。

如果 middleIndex == 0 ，左边部分的和定义为 0 。类似的，如果 middleIndex == nums.length - 1 ，右边部分的和定义为 0 。

请你返回满足上述条件 最左边 的 middleIndex ，如果不存在这样的中间位置，请你返回 -1 。
*/

/*
1. 先计算整个数组的和
2. 从左至右一次遍历，当遍历到i时， 判断i左边的和是否和右边的和相等
*/
func findMiddleIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	left_sum := 0

	for i := 0; i < len(nums); i++ {
		if sum-left_sum-nums[i] == left_sum {
			return i
		}
		left_sum += nums[i]
	}

	return -1

}

func main() {
	input := []int{2, 3, -1, 8, 4}
	result := findMiddleIndex(input)
	fmt.Printf("result = %v", result)
}
