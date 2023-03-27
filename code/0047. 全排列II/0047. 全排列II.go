package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/permutations-ii/
/*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
*/

/*
	依然采用回溯法

但在回溯的每一层中，记录每一层出现的数字，如果这个数字在该层已经出现过了，就标记为true，如果每一层都为层，则到叶子节点时，认为该数字是出现过的，不再加入
*/
func permuteUnique047(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}

func main() {
	input := []int{1, 1, 2}
	result := permuteUnique047(input)
	fmt.Printf("result = %v", result)
}
