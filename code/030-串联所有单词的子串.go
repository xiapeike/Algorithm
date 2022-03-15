package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/
/*
给定一个字符串s和一些 长度相同 的单词words 。找出 s 中恰好可以由words 中所有单词串联形成的子串的起始位置。

注意子串要与words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑words中单词串联的顺序。

*/

func findFlag(flag_map map[string]int) bool {
	flag := true
	for _, v := range flag_map {
		if v > 0 {
			flag = false
			break
		}
	}
	return flag
}

func reset_flag_map(flag_map map[string]int, words []string) map[string]int {
	for k, _ := range flag_map {
		flag_map[k] = 0
	}
	for _, word := range words {
		flag_map[word] = flag_map[word] + 1
	}
	return flag_map
}

// 使用滑动窗口暴力法，可以进行剪枝降低时间复杂度
// 思路：长度为len(words)的滑动窗口，每次后移一格，判断该字符串是否在words中，且不重复
func findSubstring(s string, words []string) []int {
	returnList := []int{}
	flag_map := make(map[string]int)

	lenth := len(words[0])
	total := len(words) * len(words[0])
	// 遍历words字符串，将flag_map进行初始化
	for _, word := range words {
		flag_map[word] = flag_map[word] + 1
	}

	// 采用滑动窗口遍历主串，若滑动窗口字符串在words数组内，且flag_map[word]=false，则滑动窗口
	index := 0 // 定义循环的位置

	for ; index <= len(s)-total; index++ {
		// 定义滑动窗口的左右边界
		left := index
		right := index + lenth
		for right <= len(s) {
			allflag := findFlag(flag_map)
			if allflag {
				returnList = append(returnList, index)
				flag_map = reset_flag_map(flag_map, words)
				break
			}
			target_string := s[left:right] // 窗口内字符串
			flag, ok := flag_map[target_string]
			if ok {
				if flag == 0 {
					break
				} else {
					flag_map[target_string] = flag_map[target_string] - 1
					left = right
					right = left + lenth
				}
			} else {
				flag_map = reset_flag_map(flag_map, words)
				break
			}
		}
		allflag := findFlag(flag_map)
		if allflag {
			returnList = append(returnList, index)
		}
		flag_map = reset_flag_map(flag_map, words)

	}
	return returnList
}

func main() {
	str := "barfoofoobarthefoobarman"
	list := []string{"bar", "foo", "the"}
	target := findSubstring(str, list)
	fmt.Printf("target = %v", target)
}
