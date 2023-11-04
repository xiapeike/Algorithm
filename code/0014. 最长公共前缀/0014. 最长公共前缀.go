package main

/*
https://leetcode.cn/problems/longest-common-prefix/?envType=study-plan-v2&envId=top-interview-150
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。
*/
import "fmt"

func longestCommonPrefix(strs []string) string {
	j := 0
	if len(strs) == 1 {
		return strs[0]
	}
out:
	for {
		for i := 0; i < len(strs)-1; i++ {
			if len(strs[i]) <= j || len(strs[i+1]) <= j {
				break out
			}
			if strs[i][j] != strs[i+1][j] {
				break out
			}
		}
		j++
	}
	if j == 0 {
		return ""
	}

	return strs[0][:j]

}

/*
当字符串数组长度为 0 时则公共前缀为空，直接返回；
令最长公共前缀 ans 的值为第一个字符串，进行初始化；
遍历后面的字符串，依次将其与 ans 进行比较，两两找出公共前缀，最终结果即为最长公共前缀；
如果查找过程中出现了 ans 为空的情况，则公共前缀不存在直接返回；
时间复杂度：O(s)O(s)O(s)，s 为所有字符串的长度之和。

作者：画手大鹏
链接：https://leetcode.cn/problems/longest-common-prefix/solutions/8348/hua-jie-suan-fa-14-zui-chang-gong-gong-qian-zhui-b/
*/
//func longestCommonPrefix2(strs []string) string {
//
//}

func main() {
	input := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(input)
	fmt.Printf("result = %v", result)
}
