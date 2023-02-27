package main

import "fmt"

//https://leetcode-cn.com/problems/longest-valid-parentheses/

/*
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
*/

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
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

/*暴力法
遍历字符串第i个（1~n-1）个字符
	遍历j（i+1~n）
		如果i-j是有效括号，则记录长度，比判断该长度是否是最大长度max，（该步骤可以剪枝，如右括号大于左括号，可提前退出该层循环）
返回最大长度
*/

func longestValidParentheses(s string) int {
	max_len := 0
	if len(s) == 0 || len(s) == 1 {
		return max_len
	}
	digits := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		digits[i] = make([]int, len(s))
	}
	// 初始化digits数组
	for i := 0; i < len(s); i++ {
		count := 0
		for j := i; j < len(s); j++ {
			if s[j] == '(' {
				count++
			} else {
				count--
			}
			digits[i][j] = count
		}

	}

	// 1、遍历字符串第i个（1~n-1）个字符
	for i := 0; i < len(s)-1; i++ {
		for j := i + max_len; j < len(s); j++ {
			str_len := j - i + 1
			//if str_len > max_len{
			if digits[i][j] < 0 {
				break
			}
			flag := isValid(s[i : j+1])
			if flag {

				if str_len > max_len {
					max_len = str_len
				}
			}
			//}

		}
	}
	return max_len
}

/*
首先，我们定义一个 dp 数组，其中第 i 个元素表示以下标为 i 的字符结尾的最长有效子字符串的长度。
我们用 dp[i] 表示以 i 结尾的最长有效括号；

（1）当 s[i] 为 (,dp[i] 必然等于 0，因为不可能组成有效的括号；

（2）当那么 s[i] 为 )

（2.1） 当 s[i-1] 为 (，那么 dp[i] = dp[i-2] + 2；

（2.2） 当 s[i-1] 为 ) 并且 s[i-dp[i-1] - 1] 为 (，那么 dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]；

时间复杂度：O(n)。
*/

func longestValidParentheses2(s string) int {
	maxAns := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			maxAns = max(maxAns, dp[i])
		}
	}
	return maxAns
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	input := "(()"
	target := longestValidParentheses2(input)
	fmt.Printf("target = %v", target)
}
