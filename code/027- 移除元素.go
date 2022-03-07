package main

import "fmt"

//https://leetcode-cn.com/problems/remove-element/
/*
给你一个数组 nums和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

*/

// 使用双指针法
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 0 // 用于控制当前插入的位置
	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			nums[index] = nums[i]
			index++
		}
	}
	fmt.Printf("result = %v", nums)
	return index
}

func main() {
	input := []int{1}
	result := removeElement(input, 1)
	fmt.Printf("result = %v", result)
}
