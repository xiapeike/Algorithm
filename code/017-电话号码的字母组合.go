package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
/*
给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

img:https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2021/11/09/200px-telephone-keypad2svg.png
*/
var phoneStr = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

/*遍历法*/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return make([]string, 0)
	}

	temp := make([]string, 1)
	for i := 0; i < len(digits); i++ {
		array, _ := phoneStr[digits[i]]
		result := make([]string, 0)
		for j := 0; j < len(array); j++ {
			digit := array[j]
			for k := 0; k < len(temp); k++ {
				result = append(result, temp[k]+string(digit))
			}

		}
		temp = result
	}
	return temp
}

/*回溯法*/
/*
总结
1、一个问题是否可用回溯法解决是很容易判断出来的，要考虑超时和构造回溯函数问题
2、回溯函数有几个共同点：
		a、一定有一个出口，及回溯到最末端时需要处理的程序
		b、当前回溯的进度，本例用index在字符串中的位置表示回溯的进度
		c、回溯中的每个阶段锁产生的临时变量需要保存下来
*/
var result []string

func huisu(index int, digits string, str string) {
	if index == len(digits) {
		// 出口
		result = append(result, str)
		return
	}
	strArray, _ := phoneStr[digits[index]]
	for i := 0; i < len(strArray); i++ {
		huisu(index+1, digits, str+string(strArray[i]))
	}
	return
}
func letterCombinations2(digits string) []string {
	if len(digits) == 0 {
		return make([]string, 0)
	}
	result = []string{}
	huisu(0, digits, "")
	return result
}

func main() {
	input := "2"
	//target := longestPalindrome(str)
	result := letterCombinations2(input)
	fmt.Printf("result = %v", result)
}
