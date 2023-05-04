package main

import (
	"code/util"
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/merge-intervals/
/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

*/

/*
首先，我们将列表中的区间按照左端点升序排序。然后我们将第一个区间加入 result 数组中，并按顺序依次考虑之后的每个区间：

如果当前区间的左端点在数组 result 中最后一个区间的右端点之后，那么它们不会重合，我们可以直接将这个区间加入数组 result 的末尾；

否则，它们重合，我们需要用当前区间的右端点更新数组 result 中最后一个区间的右端点，将其置为二者的较大值。
*/
func merge056(intervals [][]int) [][]int {
	result := [][]int{}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > result[len(result)-1][1] {
			result = append(result, intervals[i])
		} else {
			result[len(result)-1][1] = util.Max(result[len(result)-1][1], intervals[i][1])
		}
	}

	return result
}

/*
prev 初始为第一个区间，cur 表示当前的区间，res 表示结果数组

开启遍历，尝试合并 prev 和 cur，合并后更新到 prev
合并后的新区间还可能和后面的区间重合，继续尝试合并新的 cur，更新给 prev
直到不能合并 —— prev[1] < cur[0]，此时将 prev 区间推入 res 数组

*/

func merge0562(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] { // 没有一点重合
			res = append(res, prev)
			prev = cur
		} else { // 有重合
			prev[1] = util.Max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}

func main() {
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result := merge056(input)
	fmt.Printf("result = %v", result)
}
