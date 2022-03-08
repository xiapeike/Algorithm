package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/implement-strstr/
/*
实现strStr()函数。

给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。


*/

// 暴力解法，双重循环
func strStr(haystack string, needle string) int {
	result := -1
a:
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		index := i
		j := 0
		for ; j < len(needle); j++ {
			if haystack[index] == needle[j] {
				index++
			} else {
				break
			}

		}
		if j == len(needle) {
			result = i
			break a
		}
	}
	return result
}

// 滑动窗口法
func strStr2(haystack string, needle string) int {
	// 定义一个窗口两端指针
	start := 0
	end := start + len(needle)
	result := -1
	// 滑动窗口
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		str := haystack[start:end]
		if str == needle {
			result = start
			break
		}
		start++
		end++

	}
	return result
}

// KMP算法：KMP 算法是一个快速查找匹配串的算法，它的作用其实就是本题问题：如何快速在「原字符串」中找到「匹配字符串」。复杂度为 O(m + n)
// KMP 之所以能够在 O(m + n) 复杂度内完成查找，是因为其能在「非完全匹配」的过程中提取到有效信息进行复用，以减少「重复匹配」的消耗。及所谓剪枝
/*

 */
func strStr3(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

func main() {
	str := "abaaa"
	target := strStr2(str, "aa")
	fmt.Printf("sum = %v", target)
}
