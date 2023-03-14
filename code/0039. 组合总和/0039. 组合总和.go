package main

import (
	"fmt"
)

// https://leetcode.cn/problems/combination-sum/
/*
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为 target 的不同组合数少于 150 个。
*/

func combinationSum0039(candidates []int, target int) [][]int {
	var result [][]int
	result = [][]int{}
	// 递归出口
	if target == 0 {

		return [][]int{}
	}
	if len(candidates) == 0 {
		//if target%candidates[0] == 0 {
		//	var array []int
		//	for i := 0; i < target/candidates[0]; i++ {
		//		array = append(array, candidates[0])
		//	}
		//	result = append(result, array)
		//	return result
		//}
		return result
	}
	for i := 0; i < len(candidates); i++ {
		var arrays [][]int

		for j := 1; j <= target/candidates[i]; j++ {
			if target%candidates[i] == 0 && j == target/candidates[i] {
				var item []int
				for j := 1; j <= target/candidates[i]; j++ {
					item = append(item, candidates[i])
				}
				result = append(result, item)
				continue
			}

			arrays = combinationSum0039(candidates[i+1:], target-j*candidates[i])

			for _, array := range arrays {
				var item []int
				for k := 0; k < j; k++ {
					item = append(item, candidates[i])
				}
				item = append(item, array...)
				result = append(result, item)
			}
		}

	}
	return result
}

func combinationSum(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}

func main() {
	input := []int{2, 3, 5}
	result := combinationSum0039(input, 8)
	fmt.Printf("result = %v", result)
}
