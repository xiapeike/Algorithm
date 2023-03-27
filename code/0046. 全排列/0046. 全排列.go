package main

import (
	"fmt"
)

// https://leetcode.cn/problems/permutations/
/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
*/

/*
回溯法
*/
func permuteUnique047(nums []int) [][]int {
	result := [][]int{}

	var dfs func(int, int)
	var cur []int

	flag_map := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		flag_map[nums[i]] = false
	}

	dfs = func(num int, cnt int) {
		// 回溯的出口
		if cnt == len(nums) {
			result = append(result, append([]int{}, cur...))
			return
		}

		for i := 0; i < len(nums); i++ {
			flag, _ := flag_map[nums[i]]
			if flag == false {
				flag_map[nums[i]] = true
				cur = append(cur, nums[i])
				cnt++
				dfs(nums[i], cnt)
				flag_map[nums[i]] = false
				cur = cur[:len(cur)-1]
				cnt--
			}
		}

	}

	for i := 0; i < len(nums); i++ {
		flag_map[nums[i]] = true
		cur = append(cur, nums[i])
		dfs(nums[i], 1)
		flag_map[nums[i]] = false
		cur = cur[:len(cur)-1]
	}

	return result
}

func main() {
	input := []int{2}
	result := permuteUnique047(input)
	fmt.Printf("result = %v", result)
}
