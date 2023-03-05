package main

import "fmt"

// https://leetcode.cn/problems/valid-sudoku/
/*
请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）


注意：

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
空白格用 '.' 表示。
*/

// 遍历每一行，每一列，看是否元素重复
// 遍历每一个9宫格，看元素是否重复，共9个 暴力法 复杂度O(n^3)
func isValidSudoku036(board [][]byte) bool {
	uniq_map := make(map[byte]bool)
	// 遍历每一行
	for i := 0; i < len(board); i++ {
		uniq_map = make(map[byte]bool)
		for j := 0; j < len(board[i]); j++ {
			_, ok := uniq_map[board[i][j]]
			if ok == true {
				return false
			}
			if board[i][j] != '.' {
				uniq_map[board[i][j]] = true
			}

		}
	}

	// 遍历每一列
	for i := 0; i < len(board); i++ {
		uniq_map = make(map[byte]bool)
		for j := 0; j < len(board[i]); j++ {
			_, ok := uniq_map[board[j][i]]
			if ok == true {
				return false
			}
			if board[j][i] != '.' {
				uniq_map[board[j][i]] = true
			}

		}
	}

	// 遍历每个九宫格， 共9个
	for i := 0; i < len(board); i = i + 3 {
		for j := 0; j < len(board[i]); j = j + 3 {
			// 遍历每个九宫格
			uniq_map = make(map[byte]bool)
			for m := 0; m < 3; m++ {
				for n := 0; n < 3; n++ {
					_, ok := uniq_map[board[m+i][n+j]]
					if ok == true {
						return false
					}
					if board[m+i][n+j] != '.' {
						uniq_map[board[m+i][n+j]] = true
					}
				}
			}
		}
	}

	return true
}

// 解法二 添加缓存，时间复杂度 O(n^2)
/*同一个数字在每一行只能出现一次；

同一个数字在每一列只能出现一次；

同一个数字在每一个小九宫格只能出现一次。

一次定义三个数组，定义每个数字出现的次数
*/
func isValidSudoku0362(board [][]byte) bool {
	rowbuf, colbuf, boxbuf := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowbuf[i] = make([]bool, 9)
		colbuf[i] = make([]bool, 9)
		boxbuf[i] = make([]bool, 9)
	}
	// 遍历一次，添加缓存
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				num := board[r][c] - '0' - byte(1)
				if rowbuf[r][num] || colbuf[c][num] || boxbuf[r/3*3+c/3][num] {
					return false
				}
				rowbuf[r][num] = true
				colbuf[c][num] = true
				boxbuf[r/3*3+c/3][num] = true // r,c 转换到box方格中
			}
		}
	}
	return true
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
	result := isValidSudoku036(input)
	fmt.Printf("result = %v", result)
}
