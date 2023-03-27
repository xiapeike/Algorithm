package main

import (
	"fmt"
)

// https://leetcode.cn/problems/rotate-image/
/*
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
*/

/*
方法一： 借用辅助数组
每一行元素经过旋转后的位置公式：matrix[row][col] =matrix_new[col][n−row−1]
*/
func rotate048(matrix [][]int) [][]int {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}
	for i, row := range matrix {
		for j, v := range row {
			tmp[j][n-1-i] = v
		}
	}
	copy(matrix, tmp) // 拷贝 tmp 矩阵每行的引用
	return matrix
}

/*
方法二：原地旋转
每一行元素经过旋转后的位置公式：matrix[row][col] =matrix_new[col][n−row−1]
*/
func rotate0482(matrix [][]int) [][]int {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
	return matrix
}

/*
方法三：先水平翻转，再主对角线翻转
每一行元素经过旋转后的位置公式：matrix[row][col] =matrix_new[col][n−row−1]
*/
func rotate0483(matrix [][]int) [][]int {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func main() {
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	result := rotate048(input)
	fmt.Printf("result = %v", result)
}
