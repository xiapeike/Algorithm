package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/3sum/
/*
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。
*/

/*排序+双指针*/
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	// 外层遍历整个数组
	for i := 0; i < len(nums); i++ {
		num_i := nums[i]
		if num_i > 0 {
			return result
		}
		// 跳过重复值
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 内层采用双指针
		l := i + 1
		r := len(nums) - 1
		k := 0 - num_i
		for {
			if l >= r {
				break
			}
			if nums[l]+nums[r] == k {
				//add result
				result = append(result, []int{nums[i], nums[l], nums[r]})
				// 跳过重复值
				for l < r && nums[l+1] == nums[l] {
					l++
				}

				for l < r && nums[r-1] == nums[r] {
					r--
				}
				l++
				r--
			} else if nums[l]+nums[r] > k {
				// 右指针往左移，找更小的数
				r--
			} else {
				// 左指针往右移，找更大的数
				l++
			}
		}

	}
	return result
}

func main() {
	input := []int{-1, -1, 0, 0, 1}
	//target := longestPalindrome(str)
	result := threeSum(input)
	fmt.Printf("result = %v", result)
}
