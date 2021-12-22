package main

// https://leetcode-cn.com/problems/zigzag-conversion/
/*
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行Z 字形排列。

比如输入字符串为 "PAYPALISHIRING"行数为 3 时，排列如下：
P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
*/
import (
	"fmt"
)

/*
	解法一：构建一个二维数组，按顺序将字符填充，再遍历得到结果
*/
func convert(s string, numRows int) string {
	// 初始化二维字符数组
	arr := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		arr[i] = make([]byte, len(s))
	}
	// 初始化数组
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(s); j++ {
			arr[i][j] = '_'
		}
	}
	var row, col = 0, 0
	var flag = 1 // 1表示向下走， -1表示向上走
	// 遍历字符串s，按照规律写入数组
	for i := 0; i < len(s); i++ {
		cha := s[i]
		arr[row][col] = cha
		if numRows == 1 {
			col++
		} else if row == 0 {
			flag = 1
			row = row + flag
		} else if row == numRows-1 {
			flag = -1
			row = row + flag
			col++
		} else {
			row = row + flag
			if flag == -1 {
				col++
			}
		}

	}
	str := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(s); j++ {
			if arr[i][j] != '_' {
				str = str + string(arr[i][j])
			}

		}
	}
	return str
}

func main() {
	str := "PAYPALISHIRING"
	//target := longestPalindrome(str)
	target := convert(str, 3)
	fmt.Printf("sum = %v", target)
}
