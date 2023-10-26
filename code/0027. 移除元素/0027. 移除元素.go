package main

import "fmt"

/*
https://leetcode.cn/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/

/*
遍历数组，用一个单独遍历，存储目前待插入元素的位置
*/
func removeElement(nums []int, val int) int {
	count := 0
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			count++
		} else {
			if k != i {
				nums[k] = nums[i]
			}
			k++
		}

	}

	return len(nums) - count
}

func main() {
	input := []int{0, 1, 2, 2, 3, 0, 4, 2}
	result := removeElement(input, 2)
	fmt.Printf("result = %v", result)
}
