package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/count-and-say/
/*
给定一个正整数 n ，输出外观数列的第 n 项。

「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。

你可以将其视作是由递归公式定义的数字字符串序列：

countAndSay(1) = "1"
countAndSay(n) 是对 countAndSay(n-1) 的描述，然后转换成另一个数字字符串。
前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
第一项是数字 1
描述前一项，这个数是 1 即 “ 一 个 1 ”，记作 "11"
描述前一项，这个数是 11 即 “ 二 个 1 ” ，记作 "21"
描述前一项，这个数是 21 即 “ 一 个 2 + 一 个 1 ” ，记作 "1211"
描述前一项，这个数是 1211 即 “ 一 个 1 + 一 个 2 + 二 个 1 ” ，记作 "111221"
要 描述 一个数字字符串，首先要将字符串分割为 最小 数量的组，每个组都由连续的最多 相同字符 组成。然后对于每个组，先描述字符的数量，然后描述字符，形成一个描述组。要将描述转换为数字字符串，先将每组中的字符数量用数字替换，再将所有描述组连接起来。
*/

func countAndSay038(n int) string {
	if n == 1 {
		return "1"
	}
	result := ""
	last_return := countAndSay038(n - 1)
	var flag_str byte
	if len(last_return) > 0 {
		flag_str = last_return[0]
	}
	count := 0
	for i := 0; i < len(last_return); i++ {
		str := last_return[i]
		if str == flag_str {
			count = count + 1
		} else {
			result = result + strconv.Itoa(count) + string(flag_str)
			flag_str = str
			count = 1
		}
	}

	result = result + strconv.Itoa(count) + string(flag_str)

	return result
}

func countAndSay0382(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		cur := &strings.Builder{}
		for j, start := 0, 0; j < len(prev); start = j {
			for j < len(prev) && prev[j] == prev[start] {
				j++
			}
			cur.WriteString(strconv.Itoa(j - start))
			cur.WriteByte(prev[start])
		}
		prev = cur.String()
	}
	return prev
}

func countAndSay0383(n int) string {
	init_str := "1"
	i := 0

	for j := 1; j < n; j++ {
		result := &strings.Builder{}
		for k := 0; k < len(init_str); k++ {
			if init_str[i] != init_str[k] {
				result.WriteString(strconv.Itoa(k - i))
				result.WriteByte(init_str[i])
				i = k
			}
		}
		result.WriteString(strconv.Itoa(len(init_str) - i))
		result.WriteByte(init_str[i])
		init_str = result.String()
	}

	return init_str

}

func main() {
	input := 4
	result := countAndSay0383(input)
	fmt.Printf("result = %v", result)
}
