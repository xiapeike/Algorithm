package main

// https://leetcode-cn.com/problems/regular-expression-matching/
/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
*/
import (
	"fmt"
)

/*解法：穷举法*/
func isMatch(s string, p string) bool {
	len_s := len(s)
	len_p := len(p)
	i, j := 0, 0
	for i < len_s && j < len_p {
		if j+1 < len_p && p[j+1] != '*' {
			if p[j] == '.' {
				j = j + 1
				i = i + 1
			} else {
				if s[i] != p[j] {
					return false
				} else {
					j = j + 1
					i = i + 1
				}
			}
		} else if j+1 < len_p && p[j+1] == '*' {
			if j+2 >= len_p || p[j+2] != s[i] {
				if p[j] == '.' {
					if j+2 >= len_p {
						return true
					} else {
						for ; i < len_s; i++ {
							flag := isMatch(s[i:], "."+p[j+2:]) || isMatch(s[i:], p[j+2:])
							if flag == true {
								return true
							}
						}
					}

				} else {
					if s[i] != p[j] {
						j = j + 2
					} else {
						if j+2 >= len_p {
							for i < len_s && s[i] == p[j] {
								i++
							}
							j = j + 2
						} else {
							return isMatch(s[i+1:], p[j:]) || isMatch(s[i:], p[j+2:])
						}
					}
				}
			} else if p[j+2] == s[i] {
				if p[j] == '.' {
					if j+2 >= len_p {
						return true
					} else {
						for ; i < len_s; i++ {
							flag := isMatch(s[i:], "."+p[j+2:]) || isMatch(s[i:], p[j+2:])
							if flag == true {
								return true
							}
						}
					}

				} else {
					if s[i] != p[j] {
						j = j + 2
					} else {
						return isMatch(s[i+1:], p[j:]) || isMatch(s[i:], p[j+2:])
					}
				}
			}
		} else if j+1 >= len_p {
			if p[j] == '.' {
				j = j + 1
				i = i + 1
			} else {
				if s[i] != p[j] {
					return false
				} else {
					j = j + 1
					i = i + 1
				}
			}
		}
	}
	if i < len_s {
		return false
	}
	if i >= len_s && j < len_p {
		for j <= len_p-2 {
			if p[j+1] == '*' {
				j = j + 2
			} else {
				return false
			}
		}
		if j >= len_p {
			return true
		}
		return false
	}
	return true
}

/*解法：动态规划法*/
func isMatch2(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
func main() {
	input := "a"
	//target := longestPalindrome(str)
	result := isMatch(input, ".*..a*")
	fmt.Printf("result = %v", result)
}
