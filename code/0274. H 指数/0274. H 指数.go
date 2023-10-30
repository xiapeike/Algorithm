package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/h-index/?envType=study-plan-v2&envId=top-interview-150
/*
给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。

根据维基百科上 h 指数的定义：h 代表“高引用次数” ，一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文，并且每篇论文 至少 被引用 h 次。如果 h 有多种可能的值，h 指数 是其中最大的那个。
*/

/*
方法一： 暴力法
h指数最大不会超过发表的论文数

外层循环遍历发表的论文数
内存遍历h指数大于论文数的个数
*/
func hIndex(citations []int) int {
	//sort.Ints(citations)
	result := 0
	if len(citations) == 1 {
		if citations[0] > 0 {
			return 1
		} else {
			return 0
		}
	}

	for i := 1; i <= len(citations); i++ {
		count := 0
		for j := 0; j < len(citations); j++ {
			if citations[j] >= i {
				count++
			}
		}
		if count >= i {
			result = i
		}
	}

	return result
}

/*
方案二：
先排序
然后降序遍历
*/
func hIndex2(citations []int) int {
	sort.Ints(citations)
	result := 0
	if len(citations) == 1 {
		if citations[0] > 0 {
			return 1
		} else {
			return 0
		}
	}
	j := len(citations) - 1
	count := 0
	// 外层遍历文章数
	for i := len(citations); i >= 1; i-- {
		for ; j >= 0; j-- {
			if citations[j] >= i {
				count++
			} else {
				break
			}
		}
		if count >= i && i > result {
			result = i
		}
	}

	return result
}

func main() {
	input := []int{3, 0, 6, 1, 5}
	result := hIndex2(input)
	fmt.Printf("result = %v", result)
}
