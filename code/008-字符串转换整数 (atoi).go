package main

// https://leetcode-cn.com/problems/string-to-integer-atoi/
/*
请你来实现一个myAtoi(string s)函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。

函数myAtoi(string s) 的算法如下：

读入字符串并丢弃无用的前导空格
检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
如果整数数超过 32 位有符号整数范围 [−2^31, 2^31 - 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −2^31 的整数应该被固定为 −2^31 ，大于 2^31− 1 的整数应该被固定为 2^31− 1 。
返回整数作为最终结果。
*/
import (
	"fmt"
	"math"
	"strconv"
)

/*
	解法一：遍历字符串，一次解析每个字符
*/
func myAtoi(s string) int {
	i := 0
	// 去除空格可以考虑strings.Trim(s, " ")
	for ; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}
	if i >= len(s) {
		return 0
	}

	flag := 1
	if s[i] == '-' && i < len(s)-1 {
		flag = -1
		i++
	} else if s[i] == '+' && i < len(s)-1 {
		i++
	}
	result := 0
	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num, _ := strconv.Atoi(string(s[i]))
			if flag == 1 {
				if result > math.MaxInt32/10 {
					return math.MaxInt32
				} else if result == math.MaxInt32/10 && num > math.MaxInt32%10 {
					return math.MaxInt32
				}
			} else {
				if result*flag < math.MinInt32/10 {
					return math.MinInt32
				} else if result*flag == math.MinInt32/10 && num*flag < math.MinInt32%10 {
					return math.MinInt32
				}
			}
			result = result*10 + num

		} else {
			break
		}
	}
	result = result * flag
	return result
}

func main() {
	input := "-2147483647"
	//target := longestPalindrome(str)
	result := myAtoi(input)
	fmt.Printf("result = %v", math.MinInt32%10)
	fmt.Printf("result = %v", result)
}
