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
KMP算法会维持一个Next数组，该数组用于保存当模式字符p[j]与主串字符s[i]不一致时，j移动的下一个位置信息
*/
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

// 自己实现kmp算法匹配
func strStr4(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	// 第一步，构建Next匹配数组
	next := make([]int, m)
	j := 0
	next[0] = -1
	k := -1
	for j < m-1 {
		if k == -1 || needle[k] == needle[j] {
			// 当needle[k] == needle[j]时 有next[j+1] = ++k = next[j] + 1
			j = j + 1
			k = k + 1
			next[j] = k
		} else {
			// 如果不等于，说明该无公共子串，只能向前找更短的公共字串
			k = next[k]
		}
	}

	// 通过第一步找到的next数组，我们知道当haystack[i]!=needle[j]时，指针i和指针j下一步的移动
	i, j := 0, 0
	for i < n && j < m {

		if j == -1 || haystack[i] == needle[j] {
			j = j + 1
			i = i + 1

		} else {
			j = next[j]
		}

		if j == m {
			return i - j
		}

	}

	return -1
}

func main() {
	//str := "abcabeabcabd"
	str := "abcabeabcabd"
	target := strStr4(str, "abcabd")
	fmt.Printf("sum = %v", target)
}
