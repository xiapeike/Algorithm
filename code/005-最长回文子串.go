package main

// https://leetcode-cn.com/problems/longest-palindromic-substring/
import (
	"fmt"
)

// 方法一：循环遍历法
func longestPalindrome(s string) string {
	str := ""
	if len(s) == 1 {
		return s
	}
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			newstr := s[i : j+1]
			ok := ishuiwen(newstr)
			if ok && len(newstr) > len(str) {
				str = newstr
			}
		}
	}
	return str
}

// 方法二：动态规划法
func longestPalindrome1(s string) string {
	if len(s) == 1 {
		return s
	}
	// 初始化动态方程二维数组，行代表子字符串起始位置，列代表子字符串结束位置
	array := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		array[i] = make([]int, len(s))
	}

	// 定义动态方程的边界值, 只有一个字符的子串为回文子串，有两个字符的子串只有相等才是回文子串
	// 初始化对角线只有一个字符的子串，都置为1表示为回文
	for i := 0; i < len(s); i++ {
		array[i][i] = 1
	}
	// 初始化对角线只有两个字符的子串，只有相等表示为回文
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			array[i][i+1] = 1
		}
	}

	// 遍历整个数组做动态转移,需注意从长度短的到长度长的遍历
	for d := 2; d <= len(s); d++ {
		for i := 0; i < len(s); i++ {
			j := i + d - 1
			if j >= len(s) {
				break
			}

			if array[i+1][j-1] == 1 && s[i] == s[j] {
				array[i][j] = 1
			}
		}
	}
	index1, index2, cat := 0, 0, 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if array[i][j] == 1 {
				if j-i > cat {
					cat = j - i
					index1 = i
					index2 = j
				}
			}
		}
	}

	return s[index1 : index2+1]
}

func ishuiwen(s string) bool {
	if len(s) == 0 {
		return false
	}

	if len(s) == 1 {
		return true
	}

	var i, j int = 0, len(s) - 1
	for {
		if i > j {
			break
		}
		if s[i] == s[j] {
			i++
			j--
		} else if s[i] != s[j] {
			return false
		}
	}

	return true

}
func main() {
	str := "aaaaa"
	//target := longestPalindrome(str)
	target := longestPalindrome1(str)
	fmt.Printf("sum = %v", target)
}
