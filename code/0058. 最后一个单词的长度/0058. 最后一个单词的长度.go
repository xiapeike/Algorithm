package main

/*
https://leetcode.cn/problems/length-of-last-word/?envType=study-plan-v2&envId=top-interview-150
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。

单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
*/
import "fmt"

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for i >= 0 {
		if s[i] == ' ' {
			i--
		} else {
			break
		}
	}
	result := 0
	for i >= 0 {
		if s[i] == ' ' {
			break
		}
		i--
		result++
	}
	return result
}

func main() {
	input := "lengthOfLast  Word  "
	result := lengthOfLastWord(input)
	fmt.Printf("result = %v", result)
}
