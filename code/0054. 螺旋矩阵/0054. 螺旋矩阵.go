package main

import (
	"fmt"
)

// https://leetcode.cn/problems/spiral-matrix/description/
/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

*/

// 模拟法
func spiralOrder054(matrix [][]int) []int {
	result := []int{}
	m, n := len(matrix), len(matrix[0])
	visit := make([][]int, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]int, n)
	}

	visit_len := 0
	direction := "right"
	i, j := 0, 0
	for visit_len < m*n {

		switch direction {
		case "right":
			if j == n {
				// 超边界
				direction = "down"
				j--
				i++
			} else {
				// 未越界
				is_visit := visit[i][j]
				if is_visit == 1 {
					direction = "down"
					i++
					j--
				} else {
					result = append(result, matrix[i][j])
					visit_len++
					visit[i][j] = 1
					j++
				}
			}
		case "down":
			if i == m {
				// 超边界
				direction = "left"
				j--
				i--
			} else {
				// 未越界
				is_visit := visit[i][j]
				if is_visit == 1 {
					direction = "left"
					i--
					j--
				} else {
					result = append(result, matrix[i][j])
					visit_len++
					visit[i][j] = 1
					i++
				}
			}
		case "left":
			if j == -1 {
				// 超边界
				direction = "up"
				j++
				i--
			} else {
				// 未越界
				is_visit := visit[i][j]
				if is_visit == 1 {
					direction = "up"
					j++
					i--
				} else {
					result = append(result, matrix[i][j])
					visit_len++
					visit[i][j] = 1
					j--
				}
			}
		case "up":
			if i == -1 {
				// 超边界
				direction = "right"
				j++
				i++
			} else {
				// 未越界
				is_visit := visit[i][j]
				if is_visit == 1 {
					direction = "right"
					i++
					j++
				} else {
					result = append(result, matrix[i][j])
					visit_len++
					visit[i][j] = 1
					i--
				}
			}
		default:
		}
	}

	return result
}

// 模拟法优化后的代码
func spiralOrder0542(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	var (
		total          = rows * columns
		order          = make([]int, total)
		row, column    = 0, 0
		directions     = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
		directionIndex = 0
	)

	for i := 0; i < total; i++ {
		order[i] = matrix[row][column]
		visited[row][column] = true
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}
	return order
}

/*
重写
*/
func spiralOrder0543(matrix [][]int) []int {
	var result []int
	direct := [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
	direct_index := 0
	m, n := len(matrix), len(matrix[0])
	visit := make([][]int, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]int, n)
		for j := 0; j < n; j++ {
			visit[i][j] = 0
		}
	}
	x, y := 0, 0
	for i := 0; i < m*n; i++ {
		result = append(result, matrix[x][y])
		visit[x][y] = 1
		new_x, new_y := x+direct[direct_index][0], y+direct[direct_index][1]
		if new_x < 0 || new_x >= m || new_y < 0 || new_y >= n || visit[new_x][new_y] == 1 {
			direct_index = (direct_index + 1) % len(direct)
			new_x, new_y = x+direct[direct_index][0], y+direct[direct_index][1]
		}
		x, y = new_x, new_y

	}

	return result

}
func main() {
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	result := spiralOrder0543(input)
	fmt.Printf("result = %v", result)
}
