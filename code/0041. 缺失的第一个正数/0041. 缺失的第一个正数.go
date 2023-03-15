package main

import (
	"fmt"
)

// https://leetcode.cn/problems/first-missing-positive/
/*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
*/
/*
时间复杂度O(n), 空间复杂度O(n)
*/
func firstMissingPositive041(nums []int) int {
	result := 1
	result_map := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if nums[i] < result {
			continue
		} else if nums[i] > result {
			result_map[nums[i]] = true
		} else {
			j := result + 1
			for {
				if _, ok := result_map[j]; ok == false {
					result = j
					break
				}
				j++

			}
		}
	}
	return result
}

/*
循环数组做以下操作。

如果nums[i]为i+1的话，什么都不做，说明这个正数已经在正确的位置。
如果nums[i]<=0或者>len(nums)的话，说明此值已经越界，不在考虑范围，内容清零。
如果nums[i]在[1,len(nums)]的范围内(可能满足第二步，skip)， 取出此值val，内容清零。
移动到val-1的数组位置，备份val-1位置的值， 设定val-1位置的值val，目的是让val的正数归位。根据备份的值，继续第五步。 终止条件为，某位置的值<=0或者>len(nums)或者等于所在index+1的值。
循环数组，某位置为0，则返回位置index+1。如果全部正数归位，则len(nums)+1

原理：对于一个长度为 N 的数组，其中没有出现的最小正整数只能在 [1,N+1] 中。这是因为如果 [1,N] 都出现了，那么答案是 N+1，否则答案是 [1,N] 中没有出现的最小正整数
*/

func firstMissingPositive0412(nums []int) int {
	for i := 0; i < len(nums); i++ {
		s := nums[i]
		if s > len(nums) || s <= 0 {
			nums[i] = 0
		} else if i+1 != s {
			t := s
			nums[i] = 0
			for {
				b := nums[t-1]
				nums[t-1] = t
				if b > len(nums) || b <= 0 || nums[b-1] == b {
					break
				}
				t = b
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	input := []int{7, 8, 9, 11, 12}
	result := firstMissingPositive0412(input)
	fmt.Printf("result = %v", result)
}
