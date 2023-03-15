package main

import (
	"code/util"
	"fmt"
)

// https://leetcode.cn/problems/trapping-rain-water/description/
/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/
/*
计算盛满水后，每一列能够装水的量，然后将每一列装水的量加起来

由于第一列和最后一列一定不能装水，因此只遍历中间n-2个，及下标[1, n-2]

对于每一列，如何计算装水量呢，只要知道该列左边最大值，和该列右边的最大值，取二者最小再减去该列的高度及为该列的盛水量
*/
func trap042(height []int) int {
	l := len(height)
	res := 0
	// 起始位置是第一个，而不是0
	// 结束位置是倒数第二个，而不是倒数第一个
	for i := 1; i < l-1; i++ {
		leftMax := 0
		// 计算左边的最高一列，包含自己
		for k := 0; k <= i; k++ {
			leftMax = util.Max(height[k], leftMax)
		}
		rightMax := 0
		// 计算右边的最高一列，包含自己
		for j := i; j < l; j++ {
			rightMax = util.Max(height[j], rightMax)
		}
		// 取二者最小值
		minVal := util.Min(leftMax, rightMax)
		water := minVal - height[i]
		res = res + water
	}
	return res
}

/*
方法一中，需要重复计算每一列左右两边的最大值，那么我们是否能够提前计算好存下来呢，及用空间换时间
定义一个数组 db_left[n]int, n为height数组的长度，该数组每一个值存的左边高度的最大值
db_right[n]int, n为height数组的长度，该数组每一个值存的右边高度的最大值
如果已知第i列的左边最大值，第i+1列能否很快求出呢，答案是可以的：
对于第i+1列左边的最大值，等于第i列左边最大值和自己本身取最大

如果已知第i列的右边边最大值，第i-1列能否很快求出呢，答案是可以的：
对于第i列右边的最大值，等于第i-1列右边最大值和自己本身取最大
*/
func trap0422(height []int) int {
	n := len(height)

	// 计算从左往右的最大高度和从右往左的最大高度
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = util.Max(leftMax[i-1], height[i])
	}
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = util.Max(rightMax[i+1], height[i])
	}
	res := 0
	// 计算每个柱子能接的雨水量
	for i := 1; i < n-1; i++ {
		h := util.Min(leftMax[i], rightMax[i]) - height[i]
		if h > 0 {
			res += h
		}
	}
	return res
}

func main() {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	result := trap042(input)
	fmt.Printf("result = %v", result)
}
