package main

import (
	"fmt"
)

// https://leetcode.cn/problems/sudoku-solver/
/*
编写一个程序，通过填充空格来解决数独问题。

数独的解法需 遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
数独部分空格内已填入了数字，空白格用 '.' 表示。
*/

// 回溯法，遍历每个空格，从1-9进行遍历填充
func solveSudoku037(board [][]byte) {
	var line, column [9][9]bool
	var block [3][3][9]bool
	var spaces [][2]int

	// 初始化数组
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				spaces = append(spaces, [2]int{i, j})
				continue
			}
			digit := board[i][j] - '1'
			line[i][digit] = true
			column[j][digit] = true
			block[i/3][j/3][digit] = true
		}
	}
	//  从第一个空字符位置进行回溯
	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}

		i, j := spaces[pos][0], spaces[pos][1]

		//flag := false
		// 遍历每个数字 1,9 ，看放在该空格是否合适
		for k := 0; k < 9; k++ {
			if !line[i][k] && !column[j][k] && !block[i/3][j/3][k] {
				line[i][k] = true
				column[j][k] = true
				block[i/3][j/3][k] = true
				board[i][j] = byte(k) + '1'
				flag := dfs(pos + 1)
				if flag == true {
					return true
				} else {
					// 回溯失败时要重置
					line[i][k] = false
					column[j][k] = false
					block[i/3][j/3][k] = false
				}
			}
		}
		return false

	}
	dfs(0)
	fmt.Printf("board = %v", board)
}

func main() {
	input := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku037(input)

}
