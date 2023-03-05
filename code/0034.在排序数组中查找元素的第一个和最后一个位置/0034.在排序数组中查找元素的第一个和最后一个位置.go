package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
/*
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
*/

func findLeftBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right-left)/2 + left
		if target == nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}
	return left
}

func findRightBound(nums []int, target int) int {
	end := len(nums) - 1
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right-left)/2 + left
		if target == nums[mid] {
			left = mid + 1
			end = mid
		} else if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}
	return end
}

func searchRange34(nums []int, target int) []int {
	start := findLeftBound(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := findRightBound(nums, target)

	return []int{start, end}
}

func searchRange34_2(nums []int, target int) []int {
	// sort.SearchInts 为 go 实现的二分查找
	start := sort.SearchInts(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := sort.SearchInts(nums, target+1) - 1
	return []int{start, end}
}

func main() {
	input := []int{2, 2, 3}
	tarket := 2
	result := searchRange34(input, tarket)
	fmt.Printf("result = %v", result)
}
