package main

import "fmt"

// https://leetcode-cn.com/problems/generate-parentheses/
/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
*/
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0

}

// 暴力法：遍历所有可能，然后判断每个字符串是否符合有效括号表达式
func generateParenthesis(n int) []string {
	var results []string
	// 使用递归生成所有 2^{2n} 个 '(' 和 ')' 字符构成的序列
	var generateAllParentheses func(n int, str string)
	generateAllParentheses = func(n int, str string) {
		if n == 0 {
			if isValid(str) {
				results = append(results, str)
			}
		} else {
			generateAllParentheses(n-1, str+"(")
			generateAllParentheses(n-1, str+")")
		}
	}
	generateAllParentheses(2*n, "")
	return results
}

// 暴力法+剪枝优化：当遍历所有子串时，如果发送右边括号比左边括号多，则剪枝
func generateParenthesis2(n int) []string {
	var results []string
	// 使用递归生成所有 2^{2n} 个 '(' 和 ')' 字符构成的序列
	var generateAllParentheses func(n int, str string, left_num int, right_num int)
	generateAllParentheses = func(n int, str string, left_num int, right_num int) {
		if n == 0 {
			if isValid(str) {
				results = append(results, str)
			}
		} else {
			if left_num >= right_num {
				generateAllParentheses(n-1, str+"(", left_num+1, right_num)
				generateAllParentheses(n-1, str+")", left_num, right_num+1)
			}
		}
	}
	generateAllParentheses(2*n, "", 0, 0)
	return results
}

// 动态规划法：当知道所有 i<n 的情况时，我们可以通过某种算法算出 i=n 的情况。
//本题最核心的思想是，考虑 i=n 时相比 n-1 组括号增加的那一组括号的位置。
func generateParenthesis3(n int) []string {
	result := make([][]string, n+1)
	result[0] = []string{""}
	for i := 1; i <= n; i++ {
		result[i] = make([]string, 0)
		for j := 0; j < i; j++ {
			for _, intside := range result[j] {
				for _, outside := range result[i-j-1] {
					result[i] = append(result[i], "("+intside+")"+outside)
				}
			}
		}
	}
	return result[n]
}

func main() {
	input := 4
	result := generateParenthesis3(input)
	fmt.Printf("result = %v", result)
}
