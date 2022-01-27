package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/valid-parentheses/
/*
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

*/
var stack []byte

func is_match_kuohao(bty1, bty2 byte) bool {
	if bty1 == '(' && bty2 != ')' {
		return false
	}
	if bty1 == '[' && bty2 != ']' {
		return false
	}
	if bty1 == '{' && bty2 != '}' {
		return false
	}
	return true
}

/*自定义栈结构，有内存泄漏风险*/
func isValid(s string) bool {
	flag := false
	if len(s) < 1 {
		return flag
	}
	for i := 0; i < len(s); i++ {
		str := s[i]
		if str == '(' || str == '{' || str == '[' {
			// 如果是左括号进栈
			stack = append(stack, str)
		} else {
			left_symbol := stack[len(stack)-1]
			flag = is_match_kuohao(left_symbol, str)
			if flag {
				stack = stack[:len(stack)-1]
			} else {
				return flag
			}

		}
	}

	if len(stack) >= 1 {
		return false
	}
	return flag
}

func isValid2(s string) bool {
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

func main() {
	input := "{{}}"
	result := isValid(input)
	fmt.Printf("result = %v", result)
}
