package main

import (
	"fmt"
)

// https://leetcode.cn/problems/word-search/
/*
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
*/

/*
回溯法

*/

func exist079(board [][]byte, word string) bool {
	result := false
	m := len(board)
	n := len(board[0])

	visit := make([][]int, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]int, n)
		for j := 0; j < n; j++ {
			visit[i][j] = 0
		}
	}

	var dps func(int, int, int)
	dps = func(row int, col int, cnt int) {
		if board[row][col] != word[cnt] {
			return
		}

		if cnt == len(word)-1 {
			result = true
			return
		}

		visit[row][col] = 1

		// 深搜 上 下 左 右
		if row > 0 && visit[row-1][col] == 0 {
			// 上
			dps(row-1, col, cnt+1)
		}

		if row < m-1 && visit[row+1][col] == 0 {
			// 下
			dps(row+1, col, cnt+1)
		}

		if col > 0 && visit[row][col-1] == 0 {
			// 左
			dps(row, col-1, cnt+1)
		}

		if col < n-1 && visit[row][col+1] == 0 {
			// 右
			dps(row, col+1, cnt+1)
		}

		visit[row][col] = 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dps(i, j, 0)
		}
	}

	return result
}

func main() {
	input := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	//input := [][]byte{{'A'}}
	//word := "A"

	//input1 := "AB"
	result := exist079(input, word)
	fmt.Printf("result = %v", result)
}
