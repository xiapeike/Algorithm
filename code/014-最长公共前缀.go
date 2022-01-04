package main

import "fmt"

// https://leetcode-cn.com/problems/longest-common-prefix/
/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

*/

//纵向扫描时，从前往后遍历所有字符串的每一列，比较相同列上的字符是否相同，如果相同则继续对下一列进行比较，如果不相同则当前列不再属于公共前缀，
//当前列之前的部分为最长公共前缀。
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main() {
	input := []string{"flower", "flow", "flight"}
	//target := longestPalindrome(str)
	result := longestCommonPrefix(input)
	fmt.Printf("result = %v", result)
}
