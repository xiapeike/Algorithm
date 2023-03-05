package main

import "fmt"

// https://leetcode.cn/problems/search-insert-position/
/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。
*/

// 使用二分法
func searchInsert035(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left > right {
		return left
	} else {
		return right
	}

}

func main() {
	input := []int{8}
	target := 7
	result := searchInsert035(input, target)
	fmt.Printf("result = %v", result)
}
