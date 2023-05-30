package main

import (
	"code/util"
	"fmt"
)

// https://leetcode.cn/problems/maximal-rectangle/
/*
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
*/

/*
方法一：暴力法，遍历每个点，计算所有已该点为左顶点的矩形面积（会超时）
方法二：
	1. 计算出矩阵的每个元素的左边连续1的数量，使用二维数组left记录，其中 left[i][j] 为矩阵第i行第j列元素的左边连续1的数量。
	2. 对于矩阵中任意一个点，我们枚举以该点为右下角的全 1 矩形。（使用单调栈，类似求解柱状图中的最大矩形）
		具体而言，当考察以 matrix[i][j]为右下角的矩形时，我们枚举满足 0≤k≤i 的所有可能的 k，此时矩阵的最大宽度就为
		left[i][j],left[i−1][j],…,left[k][j]的最小值。
*/

func maximalRectangle085(matrix [][]byte) int {
	max := 0
	row, col := len(matrix), len(matrix[0])
	left := make([][]int, row)
	for i := 0; i < row; i++ {
		left[i] = make([]int, col)
		for j := 0; j < col; j++ {
			left[i][j] = 0
		}
	}
	// 1. 计算left矩阵
	// a. 计算left矩阵第一列
	for i := 0; i < row; i++ {
		if matrix[i][0] == '1' {
			left[i][0] = 1
		}
	}
	// b. 动态规划计算left矩阵
	for i := 0; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == '1' {
				left[i][j] = left[i][j-1] + 1
			} else {
				left[i][j] = 0
			}
		}
	}

	// 2. 遍历整个left矩阵，计算已i，j为右下角的矩阵面积
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if left[i][j] > 0 {
				min_k := left[i][j]
				for k := i; k >= 0; k-- {
					length := i - k + 1
					width := util.Min(min_k, left[k][j])
					if width*length > max {
						max = width * length
					}
					min_k = width
				}
			}
		}
	}

	return max
}

func main() {
	input := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '0', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	//input := []int{2, 4}
	//input1 := "AB"
	result := maximalRectangle085(input)
	fmt.Printf("result = %v", result)
}
