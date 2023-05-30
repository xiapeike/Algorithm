package main

import (
	"fmt"
)

// https://leetcode.cn/problems/subsets/
/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。


*/

/*
思路 1
单看每个元素，都有两种选择：选入子集，或不选入子集。

比如[1,2,3]，先看1，选1或不选1，都会再看2，选2或不选2，以此类推。

考察当前枚举的数，基于选它而继续，是一个递归分支；基于不选它而继续，又是一个分支。

用索引index代表当前递归考察的数nums[index]。

当index越界时，说明所有数字考察完了，得到一个解，把它加入解集，结束当前递归分支。
*/

func subsets078(nums []int) [][]int {
	res := [][]int{}

	var dfs func(i int, list []int)
	dfs = func(i int, list []int) {
		if i == len(nums) {
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp)
			return
		}
		list = append(list, nums[i])
		dfs(i+1, list)
		list = list[:len(list)-1]
		dfs(i+1, list)
	}
	dfs(0, []int{})

	return res

}

func main() {
	input := []int{1, 2, 3}
	//input1 := "AB"
	result := subsets078(input)
	fmt.Printf("result = %v", result)
}
