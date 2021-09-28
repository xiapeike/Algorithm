package main

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
import (
	"fmt"
)

// 方法一：使用滑动窗口，关键点，滑动窗口内一定不含重复字符
func lengthOfLongestSubstring(s string) int {
	var maxLen int = 0
	var mapList = make(map[byte]int)
	var i, j int = 0, 0

	for i = 0; i < len(s); i++ {
		for ; j < len(s); j++ {
			_, ok := mapList[s[j]]
			if ok == false {
				mapList[s[j]] = j
			} else {
				lenth := j - i
				if lenth > maxLen {
					maxLen = lenth
				}
				delete(mapList, s[i])
				break
			}
		}
		if j == len(s) {
			lenth := j - i
			if lenth > maxLen {
				maxLen = lenth
			}
		}
	}
	return maxLen
}

func main() {
	str := " "
	target := lengthOfLongestSubstring(str)
	fmt.Printf("sum = %v", target)
}
