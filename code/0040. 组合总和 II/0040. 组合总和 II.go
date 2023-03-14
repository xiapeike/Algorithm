package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/combination-sum-ii/
/*
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用 一次 。

注意：解集不能包含重复的组合。
*/
/*
1. 先找到所有的组合
2. 组合去重
*/
func combinationSum2040(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	sort.Ints(candidates)
	var dfs func(target, idx int)

	dfs = func(target, idx int) {
		if target < 0 {
			return
		}

		if target == 0 {
			// 插入之前先去重
			if unique040(ans, append([]int(nil), comb...)) {
				ans = append(ans, append([]int{}, comb...))
			}
			return
		}
		if idx == len(candidates) {
			return
		}
		for i := idx; i < len(candidates); i++ {
			if candidates[i] > target { // 剪枝，提前返回
				break
			}
			if i != idx && candidates[i] == candidates[i-1] { // 同层去重
				continue
			}

			comb = append(comb, candidates[i])
			dfs(target-candidates[i], i+1)
			comb = comb[:len(comb)-1]
		}

	}
	dfs(target, 0)

	return ans
}
func unique040(ans [][]int, comb []int) bool {
	sort.Ints(comb)
	var flag bool
	flag = true
	for _, item := range ans {
		sort.Ints(item)
		if len(comb) != len(item) {
			continue
		}

		for i := 0; i < len(comb); i++ {
			if comb[i] != item[i] {
				break
			}
			if i == len(comb)-1 {
				flag = false
			}
		}
	}
	return flag
}

func combinationSum20402(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var freq [][2]int
	for _, num := range candidates {
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var sequence []int
	var dfs func(pos, rest int)
	dfs = func(pos, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), sequence...))
			return
		}
		if pos == len(freq) || rest < freq[pos][0] {
			return
		}

		dfs(pos+1, rest)

		most := min(rest/freq[pos][0], freq[pos][1])
		for i := 1; i <= most; i++ {
			sequence = append(sequence, freq[pos][0])
			dfs(pos+1, rest-i*freq[pos][0])
		}
		sequence = sequence[:len(sequence)-most]
	}
	dfs(0, target)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//input := []int{2, 5, 2, 1, 2}
	input := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	result := combinationSum2040(input, 30)
	fmt.Printf("result = %v", result)
}
