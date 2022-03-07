package main

import "fmt"

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
/*给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。

由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么nums的前 k 个元素应该保存最终结果。

将最终结果插入nums 的前 k 个位置后返回 k 。

不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
*/

// 使用双指针法
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	if len(nums) == 0 {
		return 0
	}

	index := 1 // 用于控制当前插入的位置
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}

func main() {
	input := []int{1, 1, 2}
	//target := longestPalindrome(str)
	result := removeDuplicates(input)
	fmt.Printf("result = %v", result)
}
