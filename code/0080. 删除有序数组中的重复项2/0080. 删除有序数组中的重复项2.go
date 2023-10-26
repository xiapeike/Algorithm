package main

import "fmt"

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&envId=top-interview-150
/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
*/

/*
index指针标识待插入元素的位置
用count记录每个元素重复的次数，超过2次则不更新
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	index := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			if count == 1 {
				nums[index] = nums[i]
				index++
			}
			count++
		} else {
			nums[index] = nums[i]
			index++
			count = 1
		}
	}
	return index
}

func main() {
	input := []int{1, 1, 1, 2, 2, 3}
	result := removeDuplicates(input)
	fmt.Printf("result = %v", result)
}
